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
from xcherryapi.models.local_queue_result import LocalQueueResult
from xcherryapi.models.timer_result import TimerResult
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class CommandResults(BaseModel):
    """
    CommandResults
    """ # noqa: E501
    timer_results: Optional[List[TimerResult]] = Field(default=None, alias="timerResults")
    local_queue_results: Optional[List[LocalQueueResult]] = Field(default=None, alias="localQueueResults")
    __properties: ClassVar[List[str]] = ["timerResults", "localQueueResults"]

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
        """Create an instance of CommandResults from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of each item in timer_results (list)
        _items = []
        if self.timer_results:
            for _item in self.timer_results:
                if _item:
                    _items.append(_item.to_dict())
            _dict['timerResults'] = _items
        # override the default output from pydantic by calling `to_dict()` of each item in local_queue_results (list)
        _items = []
        if self.local_queue_results:
            for _item in self.local_queue_results:
                if _item:
                    _items.append(_item.to_dict())
            _dict['localQueueResults'] = _items
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of CommandResults from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "timerResults": [TimerResult.from_dict(_item) for _item in obj.get("timerResults")] if obj.get("timerResults") is not None else None,
            "localQueueResults": [LocalQueueResult.from_dict(_item) for _item in obj.get("localQueueResults")] if obj.get("localQueueResults") is not None else None
        })
        return _obj


