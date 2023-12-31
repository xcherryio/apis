# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.3
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel
from pydantic import Field
from xcherryapi.models.command_request import CommandRequest
from xcherryapi.models.local_queue_message import LocalQueueMessage
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class AsyncStateWaitUntilResponse(BaseModel):
    """
    the output of the waitUntil API
    """ # noqa: E501
    command_request: CommandRequest = Field(alias="commandRequest")
    publish_to_local_queue: Optional[List[LocalQueueMessage]] = Field(default=None, alias="publishToLocalQueue")
    __properties: ClassVar[List[str]] = ["commandRequest", "publishToLocalQueue"]

    model_config = {
        "populate_by_name": True,
        "validate_assignment": True
    }


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Self:
        """Create an instance of AsyncStateWaitUntilResponse from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        _dict = self.model_dump(
            by_alias=True,
            exclude={
            },
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of command_request
        if self.command_request:
            _dict['commandRequest'] = self.command_request.to_dict()
        # override the default output from pydantic by calling `to_dict()` of each item in publish_to_local_queue (list)
        _items = []
        if self.publish_to_local_queue:
            for _item in self.publish_to_local_queue:
                if _item:
                    _items.append(_item.to_dict())
            _dict['publishToLocalQueue'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of AsyncStateWaitUntilResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "commandRequest": CommandRequest.from_dict(obj.get("commandRequest")) if obj.get("commandRequest") is not None else None,
            "publishToLocalQueue": [LocalQueueMessage.from_dict(_item) for _item in obj.get("publishToLocalQueue")] if obj.get("publishToLocalQueue") is not None else None
        })
        return _obj


