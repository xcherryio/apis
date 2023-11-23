# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json


from typing import Any, ClassVar, Dict, List, Optional
from pydantic import BaseModel, StrictInt
from pydantic import Field
from xcherryapi.models.app_database_config import AppDatabaseConfig
from xcherryapi.models.local_attribute_config import LocalAttributeConfig
from xcherryapi.models.process_id_reuse_policy import ProcessIdReusePolicy
try:
    from typing import Self
except ImportError:
    from typing_extensions import Self

class ProcessStartConfig(BaseModel):
    """
    ProcessStartConfig
    """ # noqa: E501
    timeout_seconds: Optional[StrictInt] = Field(default=None, alias="timeoutSeconds")
    id_reuse_policy: Optional[ProcessIdReusePolicy] = Field(default=None, alias="idReusePolicy")
    app_database_config: Optional[AppDatabaseConfig] = Field(default=None, alias="appDatabaseConfig")
    local_attribute_config: Optional[LocalAttributeConfig] = Field(default=None, alias="localAttributeConfig")
    __properties: ClassVar[List[str]] = ["timeoutSeconds", "idReusePolicy", "appDatabaseConfig", "localAttributeConfig"]

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
        """Create an instance of ProcessStartConfig from a JSON string"""
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
        # override the default output from pydantic by calling `to_dict()` of app_database_config
        if self.app_database_config:
            _dict['appDatabaseConfig'] = self.app_database_config.to_dict()
        # override the default output from pydantic by calling `to_dict()` of local_attribute_config
        if self.local_attribute_config:
            _dict['localAttributeConfig'] = self.local_attribute_config.to_dict()
        return _dict

    @classmethod
    def from_dict(cls, obj: Dict) -> Self:
        """Create an instance of ProcessStartConfig from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "timeoutSeconds": obj.get("timeoutSeconds"),
            "idReusePolicy": obj.get("idReusePolicy"),
            "appDatabaseConfig": AppDatabaseConfig.from_dict(obj.get("appDatabaseConfig")) if obj.get("appDatabaseConfig") is not None else None,
            "localAttributeConfig": LocalAttributeConfig.from_dict(obj.get("localAttributeConfig")) if obj.get("localAttributeConfig") is not None else None
        })
        return _obj


