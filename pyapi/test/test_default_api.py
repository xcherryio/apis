# coding: utf-8

"""
    xCherry APIs

    This APIs between xCherry service and SDKs

    The version of the OpenAPI document: 0.0.1
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


import unittest

from xcherryapi.api.default_api import DefaultApi


class TestDefaultApi(unittest.TestCase):
    """DefaultApi unit test stubs"""

    def setUp(self) -> None:
        self.api = DefaultApi()

    def tearDown(self) -> None:
        pass

    def test_api_v1_xcherry_service_process_execution_describe_post(self) -> None:
        """Test case for api_v1_xcherry_service_process_execution_describe_post

        describe a process execution
        """
        pass

    def test_api_v1_xcherry_service_process_execution_publish_to_local_queue_post(self) -> None:
        """Test case for api_v1_xcherry_service_process_execution_publish_to_local_queue_post

        send message(s) to be consumed within a single process execution
        """
        pass

    def test_api_v1_xcherry_service_process_execution_rpc_post(self) -> None:
        """Test case for api_v1_xcherry_service_process_execution_rpc_post

        execute a RPC method of a process execution
        """
        pass

    def test_api_v1_xcherry_service_process_execution_start_post(self) -> None:
        """Test case for api_v1_xcherry_service_process_execution_start_post

        start a process execution
        """
        pass

    def test_api_v1_xcherry_service_process_execution_stop_post(self) -> None:
        """Test case for api_v1_xcherry_service_process_execution_stop_post

        stop a process execution
        """
        pass

    def test_api_v1_xcherry_worker_async_state_execute_post(self) -> None:
        """Test case for api_v1_xcherry_worker_async_state_execute_post

        invoking AsyncState.execute API
        """
        pass

    def test_api_v1_xcherry_worker_async_state_wait_until_post(self) -> None:
        """Test case for api_v1_xcherry_worker_async_state_wait_until_post

        invoking AsyncState.waitUntil API
        """
        pass

    def test_api_v1_xcherry_worker_process_rpc_post(self) -> None:
        """Test case for api_v1_xcherry_worker_process_rpc_post

        execute a RPC method of a process execution in the worker
        """
        pass

    def test_internal_api_v1_xcherry_notify_immediate_tasks_post(self) -> None:
        """Test case for internal_api_v1_xcherry_notify_immediate_tasks_post

        for api service to tell async service that there are new immediate tasks added to the queue
        """
        pass

    def test_internal_api_v1_xcherry_notify_timer_tasks_post(self) -> None:
        """Test case for internal_api_v1_xcherry_notify_timer_tasks_post

        for api service to tell async service that there are new timer tasks added to the queue
        """
        pass


if __name__ == '__main__':
    unittest.main()