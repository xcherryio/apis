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

from xcherryapi.models.app_database_table_row_selector import AppDatabaseTableRowSelector

class TestAppDatabaseTableRowSelector(unittest.TestCase):
    """AppDatabaseTableRowSelector unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> AppDatabaseTableRowSelector:
        """Test AppDatabaseTableRowSelector
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `AppDatabaseTableRowSelector`
        """
        model = AppDatabaseTableRowSelector()
        if include_optional:
            return AppDatabaseTableRowSelector(
                primary_key = [
                    xcherryapi.models.app_database_column_value.AppDatabaseColumnValue(
                        column = '', 
                        query_value = '', )
                    ],
                initial_write = [
                    xcherryapi.models.app_database_column_value.AppDatabaseColumnValue(
                        column = '', 
                        query_value = '', )
                    ],
                conflict_mode = 'RETURN_ERROR_ON_CONFLICT'
            )
        else:
            return AppDatabaseTableRowSelector(
                primary_key = [
                    xcherryapi.models.app_database_column_value.AppDatabaseColumnValue(
                        column = '', 
                        query_value = '', )
                    ],
        )
        """

    def testAppDatabaseTableRowSelector(self):
        """Test AppDatabaseTableRowSelector"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
