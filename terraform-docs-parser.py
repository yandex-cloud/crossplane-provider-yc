#!/usr/bin/python3
import re
import sys

FIELD_REG = re.compile('\* (.+)$')
FIELD_NAME_REG = re.compile('`(.+)` - (.+)')
LEVEL_REG = re.compile('The `(.+)` block supports:')


class Field:
    def __init__(self, name, parent, description):
        self.name = name
        self.parent = parent
        self.description = description

    def __repr__(self):
        if self.parent:
            return f"parent: '{self.parent}'; name: '{self.name}'; description '{self.description}'"
        return f"name: '{self.name}'; description '{self.description}'"


def read_file(filename) -> str:
    with open(filename) as f:
        return f.read()


def write_file(filename: str, content: str):
    with open(filename, 'w') as f:
        return f.write(content)


def parse_fields(content) -> [str]:
    fields = []
    level = ''
    for line in content.split('\n'):
        m = FIELD_REG.match(line)
        if m:
            field = m.groups()[0]
            m = FIELD_NAME_REG.match(field)
            if m:
                fields.append(Field(m.group(1), level, m.group(2)))
        else:
            m = LEVEL_REG.match(line)
            if m:
                level = m.group(1)
    return fields


def insert_comment(api_content: str, descriptions: [Field]) -> str:
    result = []
    for line in api_content.split('\n'):
        if d := match_description(line, descriptions):
            result.append(f'	// {d.description}')
        result.append(line)
    return '\n'.join(result)


def match_description(line: str, descriptions: [Field]) -> Field or None:
    for d in descriptions:
        if f'tf:"{d.name}"' in line:
            return d
        if f'tf:"{d.name},' in line:
            return d
    return None


if __name__ == '__main__':
    doc_filename = sys.argv[1]
    api_filename = sys.argv[2]

    doc_content = read_file(doc_filename)
    fields = parse_fields(doc_content)

    api_content = read_file(api_filename)
    api_content = insert_comment(api_content, fields)
    write_file(api_filename, api_content)

