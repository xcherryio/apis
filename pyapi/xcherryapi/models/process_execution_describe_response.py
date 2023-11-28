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
from pydantic import BaseModel, StrictInt, StrictStr
from pydantic import Field
from xcherryapi.models.process_status import ProcessStatus
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class ProcessExecutionDescribeResponse(BaseModel):
    """
    ProcessExecutionDescribeResponse
    """ # noqa: E501
    process_execution_id: Optional[StrictStr] = Field(default=None, alias="processExecutionId")
    process_type: Optional[StrictStr] = Field(default=None, description="the process type for SDK to lookup the process definition class", alias="processType")
    worker_url: Optional[StrictStr] = Field(default=None, description="the URL for xcherry async service to make callback to worker", alias="workerUrl")
    start_timestamp: Optional[StrictInt] = Field(default=None, description="start time of the process execution", alias="startTimestamp")
    status: Optional[ProcessStatus] = None
    __properties: ClassVar[List[str]] = ["processExecutionId", "processType", "workerUrl", "startTimestamp", "status"]

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
        """Create an instance of ProcessExecutionDescribeResponse from a JSON string"""
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
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ProcessExecutionDescribeResponse from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "processExecutionId": obj.get("processExecutionId"),
            "processType": obj.get("processType"),
            "workerUrl": obj.get("workerUrl"),
            "startTimestamp": obj.get("startTimestamp"),
            "status": obj.get("status")
        })
        return _obj


