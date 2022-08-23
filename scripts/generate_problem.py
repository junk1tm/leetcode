#!/usr/bin/env python3

import sys
import unicodedata
from html import parser

import requests


class HTMLFilter(parser.HTMLParser):
    """
    Convert HTML to plain text.
    """

    def __init__(self) -> None:
        self.text = ""
        super().__init__()

    def handle_data(self, data: str) -> None:
        self.text += data

    @classmethod
    def html2text(cls, html: str) -> str:
        f = cls()
        f.feed(html)
        return f.text


GRAPHQL_QUERY = """\
query questionData($titleSlug: String!) {
    question(titleSlug: $titleSlug) {
        questionId
        title
        difficulty
        content
        codeSnippets {
            langSlug
            code
        }
    }
}
"""


def main():
    title = sys.argv[1]  # in kebab-case

    payload = {
        "operationName": "questionData",
        "variables": {"titleSlug": title},
        "query": GRAPHQL_QUERY,
    }
    r = requests.post("https://leetcode.com/graphql", json=payload)
    data = r.json()["data"]["question"]

    number = int(data["questionId"])
    html = str(data["content"])
    text = HTMLFilter.html2text(html)
    text = text.replace("\n", "\n// ")  # comment all lines
    text = unicodedata.normalize("NFKD", text)  # remove <0xa0>

    LANG_SLUG = "golang"

    code = "package leetcode\n\n"
    for snippet in list(data["codeSnippets"]):
        if snippet["langSlug"] == LANG_SLUG:
            code += str(snippet["code"])
            break

    filename = f"{number:04}_{title.replace('-', '_')}.go"
    with open(filename, "w") as f:
        # fmt: off
        f.write(f"""\
// https://leetcode.com/problems/{title}
//
// {number:04}. {data['title']} [{data['difficulty']}]
//
// {text}

{code}
""")
        # fmt: on


if __name__ == "__main__":
    main()
