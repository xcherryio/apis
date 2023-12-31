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
from xcherryapi.models.command_waiting_type import CommandWaitingType
from xcherryapi.models.local_queue_command import LocalQueueCommand
from xcherryapi.models.timer_command import TimerCommand
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class CommandRequest(BaseModel):
    """
    CommandRequest
    """ # noqa: E501
    waiting_type: CommandWaitingType = Field(alias="waitingType")
    timer_commands: Optional[List[TimerCommand]] = Field(default=None, alias="timerCommands")
    local_queue_commands: Optional[List[LocalQueueCommand]] = Field(default=None, alias="localQueueCommands")
    __properties: ClassVar[List[str]] = ["waitingType", "timerCommands", "localQueueCommands"]

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
        """Create an instance of CommandRequest from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in timer_commands (list)
        _items = []
        if self.timer_commands:
            for _item in self.timer_commands:
                if _item:
                    _items.append(_item.to_dict())
            _dict['timerCommands'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in local_queue_commands (list)
        _items = []
        if self.local_queue_commands:
            for _item in self.local_queue_commands:
                if _item:
                    _items.append(_item.to_dict())
            _dict['localQueueCommands'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of CommandRequest from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "waitingType": obj.get("waitingType"),
            "timerCommands": [TimerCommand.from_dict(_item) for _item in obj.get("timerCommands")] if obj.get("timerCommands") is not None else None,
            "localQueueCommands": [LocalQueueCommand.from_dict(_item) for _item in obj.get("localQueueCommands")] if obj.get("localQueueCommands") is not None else None
        })
        return _obj


