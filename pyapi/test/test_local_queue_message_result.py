# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from xcherryapi.models.local_queue_message_result import LocalQueueMessageResult

class TestLocalQueueMessageResult(unittest.TestCase):
    """LocalQueueMessageResult unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> LocalQueueMessageResult:
        """Test LocalQueueMessageResult
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `LocalQueueMessageResult`
        """
        model = LocalQueueMessageResult()
        if include_optional:
            return LocalQueueMessageResult(
                dedup_id = '',
                payload = xcherryapi.models.encoded_object.EncodedObject(
                    encoding = '', 
                    data = '', )
            )
        else:
            return LocalQueueMessageResult(
                dedup_id = '',
        )
        """

    def testLocalQueueMessageResult(self):
        """Test LocalQueueMessageResult"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
