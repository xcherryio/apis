# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.3
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest
import datetime

from xcherryapi.models.api_error_response import ApiErrorResponse

class TestApiErrorResponse(unittest.TestCase):
    """ApiErrorResponse unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> ApiErrorResponse:
        """Test ApiErrorResponse
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `ApiErrorResponse`
        """
        model = ApiErrorResponse()
        if include_optional:
            return ApiErrorResponse(
                error_sub_type = 'UNCATEGORIZED_ERROR',
                app_error_type = '',
                details = ''
            )
        else:
            return ApiErrorResponse(
        )
        """

    def testApiErrorResponse(self):
        """Test ApiErrorResponse"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
