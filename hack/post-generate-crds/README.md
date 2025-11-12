# CRD Post-Generation Tool

This tool addresses a known issue where `controller-gen` removes custom `x-kubernetes-validations` from CRDs during generation.

## Problem

The FolderIAMMember CRD requires an immutability validation on the `role` field in the `initProvider` section. This validation ensures that once a role is set, it cannot be changed (using CEL expression `self == oldSelf`).

However, `controller-gen` removes this validation during CRD generation, requiring manual restoration after each `make generate`.

**Reference**: https://github.com/crossplane/upjet/issues/78

## Solution

This Go tool automatically injects the required validation after CRD generation. It:

1. Reads the generated FolderIAMMember CRD
2. Locates the `role` field in the `initProvider` section
3. Injects the immutability validation before the `type: string` line
4. Writes the modified CRD back to disk

## Usage

The tool is automatically invoked as part of the code generation pipeline in `apis/generate.go`:

```go
//go:generate go run -tags generate ../hack/post-generate-crds/main.go
```

When you run `make generate`, the tool will:
- Process the FolderIAMMember CRD
- Check if the validation already exists (to avoid duplicates)
- Inject the validation if missing
- Report success or failure

## Injected Validation

The tool injects the following YAML into the CRD:

```yaml
# WARNING: This will be deleted upon generation, see https://github.com/crossplane/upjet/issues/78 -
# WARNING: restore this manually!
x-kubernetes-validations:
  - message: 'Role can not be changed after creation. If changes are needed, delete this FolderIAMMember and create a new one.'
    rule: self == oldSelf
```

## Adding More Resources

To add immutability validation to other IAM member resources, uncomment and modify the relevant section in `main.go`:

```go
// Add more IAM member CRDs here if needed in the future
if err := injectRoleImmutability(
	filepath.Join(crdDir, "iam.yandex-cloud.jet.crossplane.io_cloudiammembers.yaml"),
	"CloudIAMMember",
); err != nil {
	fmt.Fprintf(os.Stderr, "Error processing CloudIAMMember: %v\n", err)
	os.Exit(1)
}
```

## Technical Details

- **Language**: Go (to maintain consistency with the project)
- **Build Tag**: `//go:build generate` (only compiled during generation)
- **Location**: Injected at line ~371 in the CRD (after role description, before `type: string`)
- **Indentation**: Maintains proper YAML indentation (20 spaces for field level)