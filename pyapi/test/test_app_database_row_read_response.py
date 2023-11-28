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

from xcherryapi.models.app_database_row_read_response import AppDatabaseRowReadResponse

class TestAppDatabaseRowReadResponse(unittest.TestCase):
    """AppDatabaseRowReadResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AppDatabaseRowReadResponse:
        """Test AppDatabaseRowReadResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AppDatabaseRowReadResponse`
        """
        model = AppDatabaseRowReadResponse()
        if include_optional:
            return AppDatabaseRowReadResponse(
                columns = [
                    xcherryapi.models.app_database_column_value.AppDatabaseColumnValue(
                        column = '', 
                        query_value = '', )
                    ]
            )
        else:
            return AppDatabaseRowReadResponse(
        )
        """

    def testAppDatabaseRowReadResponse(self):
        """Test AppDatabaseRowReadResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
