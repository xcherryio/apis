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

from xcherryapi.models.retry_policy import RetryPolicy

class TestRetryPolicy(unittest.TestCase):
    """RetryPolicy unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def make_instance(self, include_optional) -> RetryPolicy:
        """Test RetryPolicy
            include_option is a boolean, when False only required
            params are included, when True both required and
            optional params are included """
        # uncomment below to create an instance of `RetryPolicy`
        """
        model = RetryPolicy()
        if include_optional:
            return RetryPolicy(
                initial_interval_seconds = 56,
                backoff_coefficient = 1.337,
                maximum_interval_seconds = 56,
                maximum_attempts = 56,
                maximum_attempts_duration_seconds = 56
            )
        else:
            return RetryPolicy(
        )
        """

    def testRetryPolicy(self):
        """Test RetryPolicy"""
        # inst_req_only = self.make_instance(include_optional=False)
        # inst_req_and_optional = self.make_instance(include_optional=True)

if __name__ == '__main__':
    unittest.main()
