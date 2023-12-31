# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.3
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import json
import pprint
import re  # noqa: F401
from enum import Enum



try:
    from typing import Self
except ImportError:
    from typing_extensions import Self


class ProcessIdReusePolicy(str, Enum):
    """
    ProcessIdReusePolicy
    """

    """
    allowed enum values
    """
    ALLOW_IF_PREVIOUS_EXIT_ABNORMALLY = 'ALLOW_IF_PREVIOUS_EXIT_ABNORMALLY'
    ALLOW_IF_NO_RUNNING = 'ALLOW_IF_NO_RUNNING'
    DISALLOW_REUSE = 'DISALLOW_REUSE'
    TERMINATE_IF_RUNNING = 'TERMINATE_IF_RUNNING'

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of ProcessIdReusePolicy from a JSON string"""
        return cls(json.loads(json_str))


