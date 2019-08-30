# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/service/admin.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.api import annotations_pb2 as google_dot_api_dot_annotations__pb2
from flyteidl.admin import project_pb2 as flyteidl_dot_admin_dot_project__pb2
from flyteidl.admin import task_pb2 as flyteidl_dot_admin_dot_task__pb2
from flyteidl.admin import workflow_pb2 as flyteidl_dot_admin_dot_workflow__pb2
from flyteidl.admin import launch_plan_pb2 as flyteidl_dot_admin_dot_launch__plan__pb2
from flyteidl.admin import event_pb2 as flyteidl_dot_admin_dot_event__pb2
from flyteidl.admin import execution_pb2 as flyteidl_dot_admin_dot_execution__pb2
from flyteidl.admin import node_execution_pb2 as flyteidl_dot_admin_dot_node__execution__pb2
from flyteidl.admin import task_execution_pb2 as flyteidl_dot_admin_dot_task__execution__pb2
from flyteidl.admin import common_pb2 as flyteidl_dot_admin_dot_common__pb2
from protoc_gen_swagger.options import annotations_pb2 as protoc__gen__swagger_dot_options_dot_annotations__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/service/admin.proto',
  package='flyteidl.service',
  syntax='proto3',
  serialized_pb=_b('\n\x1c\x66lyteidl/service/admin.proto\x12\x10\x66lyteidl.service\x1a\x1cgoogle/api/annotations.proto\x1a\x1c\x66lyteidl/admin/project.proto\x1a\x19\x66lyteidl/admin/task.proto\x1a\x1d\x66lyteidl/admin/workflow.proto\x1a flyteidl/admin/launch_plan.proto\x1a\x1a\x66lyteidl/admin/event.proto\x1a\x1e\x66lyteidl/admin/execution.proto\x1a#flyteidl/admin/node_execution.proto\x1a#flyteidl/admin/task_execution.proto\x1a\x1b\x66lyteidl/admin/common.proto\x1a,protoc-gen-swagger/options/annotations.proto2\x85;\n\x0c\x41\x64minService\x12\xc4\x02\n\nCreateTask\x12!.flyteidl.admin.TaskCreateRequest\x1a\".flyteidl.admin.TaskCreateResponse\"\xee\x01\x82\xd3\xe4\x93\x02\x12\"\r/api/v1/tasks:\x01*\x92\x41\xd2\x01\x1a%Create and register a task definitionJB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\x88\x01\n\x07GetTask\x12 .flyteidl.admin.ObjectGetRequest\x1a\x14.flyteidl.admin.Task\"E\x82\xd3\xe4\x93\x02?\x12=/api/v1/tasks/{id.project}/{id.domain}/{id.name}/{id.version}\x12\x97\x01\n\x0bListTaskIds\x12\x30.flyteidl.admin.NamedEntityIdentifierListRequest\x1a).flyteidl.admin.NamedEntityIdentifierList\"+\x82\xd3\xe4\x93\x02%\x12#/api/v1/task_ids/{project}/{domain}\x12\xae\x01\n\tListTasks\x12#.flyteidl.admin.ResourceListRequest\x1a\x18.flyteidl.admin.TaskList\"b\x82\xd3\xe4\x93\x02\\\x12\x30/api/v1/tasks/{id.project}/{id.domain}/{id.name}Z(\x12&/api/v1/tasks/{id.project}/{id.domain}\x12\xd8\x02\n\x0e\x43reateWorkflow\x12%.flyteidl.admin.WorkflowCreateRequest\x1a&.flyteidl.admin.WorkflowCreateResponse\"\xf6\x01\x82\xd3\xe4\x93\x02\x16\"\x11/api/v1/workflows:\x01*\x92\x41\xd6\x01\x1a)Create and register a workflow definitionJB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\x94\x01\n\x0bGetWorkflow\x12 .flyteidl.admin.ObjectGetRequest\x1a\x18.flyteidl.admin.Workflow\"I\x82\xd3\xe4\x93\x02\x43\x12\x41/api/v1/workflows/{id.project}/{id.domain}/{id.name}/{id.version}\x12\x9f\x01\n\x0fListWorkflowIds\x12\x30.flyteidl.admin.NamedEntityIdentifierListRequest\x1a).flyteidl.admin.NamedEntityIdentifierList\"/\x82\xd3\xe4\x93\x02)\x12\'/api/v1/workflow_ids/{project}/{domain}\x12\xbe\x01\n\rListWorkflows\x12#.flyteidl.admin.ResourceListRequest\x1a\x1c.flyteidl.admin.WorkflowList\"j\x82\xd3\xe4\x93\x02\x64\x12\x34/api/v1/workflows/{id.project}/{id.domain}/{id.name}Z,\x12*/api/v1/workflows/{id.project}/{id.domain}\x12\xe4\x02\n\x10\x43reateLaunchPlan\x12\'.flyteidl.admin.LaunchPlanCreateRequest\x1a(.flyteidl.admin.LaunchPlanCreateResponse\"\xfc\x01\x82\xd3\xe4\x93\x02\x19\"\x14/api/v1/launch_plans:\x01*\x92\x41\xd9\x01\x1a,Create and register a launch plan definitionJB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\x9b\x01\n\rGetLaunchPlan\x12 .flyteidl.admin.ObjectGetRequest\x1a\x1a.flyteidl.admin.LaunchPlan\"L\x82\xd3\xe4\x93\x02\x46\x12\x44/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}\x12\xa2\x01\n\x13GetActiveLaunchPlan\x12\'.flyteidl.admin.ActiveLaunchPlanRequest\x1a\x1a.flyteidl.admin.LaunchPlan\"F\x82\xd3\xe4\x93\x02@\x12>/api/v1/active_launch_plans/{id.project}/{id.domain}/{id.name}\x12\x9c\x01\n\x15ListActiveLaunchPlans\x12+.flyteidl.admin.ActiveLaunchPlanListRequest\x1a\x1e.flyteidl.admin.LaunchPlanList\"6\x82\xd3\xe4\x93\x02\x30\x12./api/v1/active_launch_plans/{project}/{domain}\x12\xa4\x01\n\x11ListLaunchPlanIds\x12\x30.flyteidl.admin.NamedEntityIdentifierListRequest\x1a).flyteidl.admin.NamedEntityIdentifierList\"2\x82\xd3\xe4\x93\x02,\x12*/api/v1/launch_plan_ids/{project}/{domain}\x12\xc8\x01\n\x0fListLaunchPlans\x12#.flyteidl.admin.ResourceListRequest\x1a\x1e.flyteidl.admin.LaunchPlanList\"p\x82\xd3\xe4\x93\x02j\x12\x37/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}Z/\x12-/api/v1/launch_plans/{id.project}/{id.domain}\x12\xb6\x01\n\x10UpdateLaunchPlan\x12\'.flyteidl.admin.LaunchPlanUpdateRequest\x1a(.flyteidl.admin.LaunchPlanUpdateResponse\"O\x82\xd3\xe4\x93\x02I\x1a\x44/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}:\x01*\x12\xce\x02\n\x0f\x43reateExecution\x12&.flyteidl.admin.ExecutionCreateRequest\x1a\'.flyteidl.admin.ExecutionCreateResponse\"\xe9\x01\x82\xd3\xe4\x93\x02\x17\"\x12/api/v1/executions:\x01*\x92\x41\xc8\x01\x1a\x1b\x43reate a workflow executionJB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\xdd\x02\n\x11RelaunchExecution\x12(.flyteidl.admin.ExecutionRelaunchRequest\x1a\'.flyteidl.admin.ExecutionCreateResponse\"\xf4\x01\x82\xd3\xe4\x93\x02 \"\x1b/api/v1/executions/relaunch:\x01*\x92\x41\xca\x01\x1a\x1dRelaunch a workflow executionJB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.Je\n\x03\x34\x30\x39\x12^\n\\Returned for a request that references an identical entity that has already been registered.\x12\x95\x01\n\x0cGetExecution\x12+.flyteidl.admin.WorkflowExecutionGetRequest\x1a\x19.flyteidl.admin.Execution\"=\x82\xd3\xe4\x93\x02\x37\x12\x35/api/v1/executions/{id.project}/{id.domain}/{id.name}\x12\xb9\x01\n\x10GetExecutionData\x12/.flyteidl.admin.WorkflowExecutionGetDataRequest\x1a\x30.flyteidl.admin.WorkflowExecutionGetDataResponse\"B\x82\xd3\xe4\x93\x02<\x12:/api/v1/data/executions/{id.project}/{id.domain}/{id.name}\x12\x89\x01\n\x0eListExecutions\x12#.flyteidl.admin.ResourceListRequest\x1a\x1d.flyteidl.admin.ExecutionList\"3\x82\xd3\xe4\x93\x02-\x12+/api/v1/executions/{id.project}/{id.domain}\x12\xad\x01\n\x12TerminateExecution\x12).flyteidl.admin.ExecutionTerminateRequest\x1a*.flyteidl.admin.ExecutionTerminateResponse\"@\x82\xd3\xe4\x93\x02:*5/api/v1/executions/{id.project}/{id.domain}/{id.name}:\x01*\x12\xd2\x01\n\x10GetNodeExecution\x12\'.flyteidl.admin.NodeExecutionGetRequest\x1a\x1d.flyteidl.admin.NodeExecution\"v\x82\xd3\xe4\x93\x02p\x12n/api/v1/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}\x12\xde\x01\n\x12ListNodeExecutions\x12(.flyteidl.admin.NodeExecutionListRequest\x1a!.flyteidl.admin.NodeExecutionList\"{\x82\xd3\xe4\x93\x02u\x12s/api/v1/node_executions/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}\x12\xa5\x04\n\x19ListNodeExecutionsForTask\x12/.flyteidl.admin.NodeExecutionForTaskListRequest\x1a!.flyteidl.admin.NodeExecutionList\"\xb3\x03\x82\xd3\xe4\x93\x02\xac\x03\x12\xa9\x03/api/v1/children/task_executions/{task_execution_id.node_execution_id.execution_id.project}/{task_execution_id.node_execution_id.execution_id.domain}/{task_execution_id.node_execution_id.execution_id.name}/{task_execution_id.node_execution_id.node_id}/{task_execution_id.task_id.project}/{task_execution_id.task_id.domain}/{task_execution_id.task_id.name}/{task_execution_id.task_id.version}/{task_execution_id.retry_attempt}\x12\xee\x01\n\x14GetNodeExecutionData\x12+.flyteidl.admin.NodeExecutionGetDataRequest\x1a,.flyteidl.admin.NodeExecutionGetDataResponse\"{\x82\xd3\xe4\x93\x02u\x12s/api/v1/data/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}\x12\xa5\x02\n\x0fRegisterProject\x12&.flyteidl.admin.ProjectRegisterRequest\x1a\'.flyteidl.admin.ProjectRegisterResponse\"\xc0\x01\x82\xd3\xe4\x93\x02\x15\"\x10/api/v1/projects:\x01*\x92\x41\xa1\x01\x1a+Register a project along with valid domainsJ.\n\x03\x32\x30\x31\x12\'\n%Returned for successful registration.JB\n\x03\x34\x30\x30\x12;\n9Returned for bad request that may have failed validation.\x12\x66\n\x0cListProjects\x12\".flyteidl.admin.ProjectListRequest\x1a\x18.flyteidl.admin.Projects\"\x18\x82\xd3\xe4\x93\x02\x12\x12\x10/api/v1/projects\x12\x99\x01\n\x13\x43reateWorkflowEvent\x12-.flyteidl.admin.WorkflowExecutionEventRequest\x1a..flyteidl.admin.WorkflowExecutionEventResponse\"#\x82\xd3\xe4\x93\x02\x1d\"\x18/api/v1/events/workflows:\x01*\x12\x89\x01\n\x0f\x43reateNodeEvent\x12).flyteidl.admin.NodeExecutionEventRequest\x1a*.flyteidl.admin.NodeExecutionEventResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/api/v1/events/nodes:\x01*\x12\x89\x01\n\x0f\x43reateTaskEvent\x12).flyteidl.admin.TaskExecutionEventRequest\x1a*.flyteidl.admin.TaskExecutionEventResponse\"\x1f\x82\xd3\xe4\x93\x02\x19\"\x14/api/v1/events/tasks:\x01*\x12\x80\x03\n\x10GetTaskExecution\x12\'.flyteidl.admin.TaskExecutionGetRequest\x1a\x1d.flyteidl.admin.TaskExecution\"\xa3\x02\x82\xd3\xe4\x93\x02\x9c\x02\x12\x99\x02/api/v1/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}\x12\x98\x02\n\x12ListTaskExecutions\x12(.flyteidl.admin.TaskExecutionListRequest\x1a!.flyteidl.admin.TaskExecutionList\"\xb4\x01\x82\xd3\xe4\x93\x02\xad\x01\x12\xaa\x01/api/v1/task_executions/{node_execution_id.execution_id.project}/{node_execution_id.execution_id.domain}/{node_execution_id.execution_id.name}/{node_execution_id.node_id}\x12\x9c\x03\n\x14GetTaskExecutionData\x12+.flyteidl.admin.TaskExecutionGetDataRequest\x1a,.flyteidl.admin.TaskExecutionGetDataResponse\"\xa8\x02\x82\xd3\xe4\x93\x02\xa1\x02\x12\x9e\x02/api/v1/data/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}B5Z3github.com/lyft/flyteidl/gen/pb-go/flyteidl/serviceb\x06proto3')
  ,
  dependencies=[google_dot_api_dot_annotations__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_project__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_task__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_workflow__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_launch__plan__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_event__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_execution__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_node__execution__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_task__execution__pb2.DESCRIPTOR,flyteidl_dot_admin_dot_common__pb2.DESCRIPTOR,protoc__gen__swagger_dot_options_dot_annotations__pb2.DESCRIPTOR,])



_sym_db.RegisterFileDescriptor(DESCRIPTOR)


DESCRIPTOR.has_options = True
DESCRIPTOR._options = _descriptor._ParseOptions(descriptor_pb2.FileOptions(), _b('Z3github.com/lyft/flyteidl/gen/pb-go/flyteidl/service'))

_ADMINSERVICE = _descriptor.ServiceDescriptor(
  name='AdminService',
  full_name='flyteidl.service.AdminService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=412,
  serialized_end=7969,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreateTask',
    full_name='flyteidl.service.AdminService.CreateTask',
    index=0,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__pb2._TASKCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_task__pb2._TASKCREATERESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\022\"\r/api/v1/tasks:\001*\222A\322\001\032%Create and register a task definitionJB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.')),
  ),
  _descriptor.MethodDescriptor(
    name='GetTask',
    full_name='flyteidl.service.AdminService.GetTask',
    index=1,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._OBJECTGETREQUEST,
    output_type=flyteidl_dot_admin_dot_task__pb2._TASK,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002?\022=/api/v1/tasks/{id.project}/{id.domain}/{id.name}/{id.version}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListTaskIds',
    full_name='flyteidl.service.AdminService.ListTaskIds',
    index=2,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002%\022#/api/v1/task_ids/{project}/{domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListTasks',
    full_name='flyteidl.service.AdminService.ListTasks',
    index=3,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_task__pb2._TASKLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\\\0220/api/v1/tasks/{id.project}/{id.domain}/{id.name}Z(\022&/api/v1/tasks/{id.project}/{id.domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='CreateWorkflow',
    full_name='flyteidl.service.AdminService.CreateWorkflow',
    index=4,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOWCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOWCREATERESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\026\"\021/api/v1/workflows:\001*\222A\326\001\032)Create and register a workflow definitionJB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.')),
  ),
  _descriptor.MethodDescriptor(
    name='GetWorkflow',
    full_name='flyteidl.service.AdminService.GetWorkflow',
    index=5,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._OBJECTGETREQUEST,
    output_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOW,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002C\022A/api/v1/workflows/{id.project}/{id.domain}/{id.name}/{id.version}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListWorkflowIds',
    full_name='flyteidl.service.AdminService.ListWorkflowIds',
    index=6,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002)\022\'/api/v1/workflow_ids/{project}/{domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListWorkflows',
    full_name='flyteidl.service.AdminService.ListWorkflows',
    index=7,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_workflow__pb2._WORKFLOWLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002d\0224/api/v1/workflows/{id.project}/{id.domain}/{id.name}Z,\022*/api/v1/workflows/{id.project}/{id.domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='CreateLaunchPlan',
    full_name='flyteidl.service.AdminService.CreateLaunchPlan',
    index=8,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANCREATERESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\031\"\024/api/v1/launch_plans:\001*\222A\331\001\032,Create and register a launch plan definitionJB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.')),
  ),
  _descriptor.MethodDescriptor(
    name='GetLaunchPlan',
    full_name='flyteidl.service.AdminService.GetLaunchPlan',
    index=9,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._OBJECTGETREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLAN,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002F\022D/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}')),
  ),
  _descriptor.MethodDescriptor(
    name='GetActiveLaunchPlan',
    full_name='flyteidl.service.AdminService.GetActiveLaunchPlan',
    index=10,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._ACTIVELAUNCHPLANREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLAN,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002@\022>/api/v1/active_launch_plans/{id.project}/{id.domain}/{id.name}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListActiveLaunchPlans',
    full_name='flyteidl.service.AdminService.ListActiveLaunchPlans',
    index=11,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._ACTIVELAUNCHPLANLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\0020\022./api/v1/active_launch_plans/{project}/{domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListLaunchPlanIds',
    full_name='flyteidl.service.AdminService.ListLaunchPlanIds',
    index=12,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_common__pb2._NAMEDENTITYIDENTIFIERLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002,\022*/api/v1/launch_plan_ids/{project}/{domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListLaunchPlans',
    full_name='flyteidl.service.AdminService.ListLaunchPlans',
    index=13,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002j\0227/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}Z/\022-/api/v1/launch_plans/{id.project}/{id.domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='UpdateLaunchPlan',
    full_name='flyteidl.service.AdminService.UpdateLaunchPlan',
    index=14,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANUPDATEREQUEST,
    output_type=flyteidl_dot_admin_dot_launch__plan__pb2._LAUNCHPLANUPDATERESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002I\032D/api/v1/launch_plans/{id.project}/{id.domain}/{id.name}/{id.version}:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='CreateExecution',
    full_name='flyteidl.service.AdminService.CreateExecution',
    index=15,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONCREATEREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONCREATERESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\027\"\022/api/v1/executions:\001*\222A\310\001\032\033Create a workflow executionJB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.')),
  ),
  _descriptor.MethodDescriptor(
    name='RelaunchExecution',
    full_name='flyteidl.service.AdminService.RelaunchExecution',
    index=16,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONRELAUNCHREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONCREATERESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002 \"\033/api/v1/executions/relaunch:\001*\222A\312\001\032\035Relaunch a workflow executionJB\n\003400\022;\n9Returned for bad request that may have failed validation.Je\n\003409\022^\n\\Returned for a request that references an identical entity that has already been registered.')),
  ),
  _descriptor.MethodDescriptor(
    name='GetExecution',
    full_name='flyteidl.service.AdminService.GetExecution',
    index=17,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._WORKFLOWEXECUTIONGETREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTION,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\0027\0225/api/v1/executions/{id.project}/{id.domain}/{id.name}')),
  ),
  _descriptor.MethodDescriptor(
    name='GetExecutionData',
    full_name='flyteidl.service.AdminService.GetExecutionData',
    index=18,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._WORKFLOWEXECUTIONGETDATAREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._WORKFLOWEXECUTIONGETDATARESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002<\022:/api/v1/data/executions/{id.project}/{id.domain}/{id.name}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListExecutions',
    full_name='flyteidl.service.AdminService.ListExecutions',
    index=19,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_common__pb2._RESOURCELISTREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002-\022+/api/v1/executions/{id.project}/{id.domain}')),
  ),
  _descriptor.MethodDescriptor(
    name='TerminateExecution',
    full_name='flyteidl.service.AdminService.TerminateExecution',
    index=20,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONTERMINATEREQUEST,
    output_type=flyteidl_dot_admin_dot_execution__pb2._EXECUTIONTERMINATERESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002:*5/api/v1/executions/{id.project}/{id.domain}/{id.name}:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='GetNodeExecution',
    full_name='flyteidl.service.AdminService.GetNodeExecution',
    index=21,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONGETREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTION,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002p\022n/api/v1/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListNodeExecutions',
    full_name='flyteidl.service.AdminService.ListNodeExecutions',
    index=22,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002u\022s/api/v1/node_executions/{workflow_execution_id.project}/{workflow_execution_id.domain}/{workflow_execution_id.name}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListNodeExecutionsForTask',
    full_name='flyteidl.service.AdminService.ListNodeExecutionsForTask',
    index=23,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONFORTASKLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\254\003\022\251\003/api/v1/children/task_executions/{task_execution_id.node_execution_id.execution_id.project}/{task_execution_id.node_execution_id.execution_id.domain}/{task_execution_id.node_execution_id.execution_id.name}/{task_execution_id.node_execution_id.node_id}/{task_execution_id.task_id.project}/{task_execution_id.task_id.domain}/{task_execution_id.task_id.name}/{task_execution_id.task_id.version}/{task_execution_id.retry_attempt}')),
  ),
  _descriptor.MethodDescriptor(
    name='GetNodeExecutionData',
    full_name='flyteidl.service.AdminService.GetNodeExecutionData',
    index=24,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONGETDATAREQUEST,
    output_type=flyteidl_dot_admin_dot_node__execution__pb2._NODEEXECUTIONGETDATARESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002u\022s/api/v1/data/node_executions/{id.execution_id.project}/{id.execution_id.domain}/{id.execution_id.name}/{id.node_id}')),
  ),
  _descriptor.MethodDescriptor(
    name='RegisterProject',
    full_name='flyteidl.service.AdminService.RegisterProject',
    index=25,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_project__pb2._PROJECTREGISTERREQUEST,
    output_type=flyteidl_dot_admin_dot_project__pb2._PROJECTREGISTERRESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\025\"\020/api/v1/projects:\001*\222A\241\001\032+Register a project along with valid domainsJ.\n\003201\022\'\n%Returned for successful registration.JB\n\003400\022;\n9Returned for bad request that may have failed validation.')),
  ),
  _descriptor.MethodDescriptor(
    name='ListProjects',
    full_name='flyteidl.service.AdminService.ListProjects',
    index=26,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_project__pb2._PROJECTLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_project__pb2._PROJECTS,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\022\022\020/api/v1/projects')),
  ),
  _descriptor.MethodDescriptor(
    name='CreateWorkflowEvent',
    full_name='flyteidl.service.AdminService.CreateWorkflowEvent',
    index=27,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_event__pb2._WORKFLOWEXECUTIONEVENTREQUEST,
    output_type=flyteidl_dot_admin_dot_event__pb2._WORKFLOWEXECUTIONEVENTRESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\035\"\030/api/v1/events/workflows:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='CreateNodeEvent',
    full_name='flyteidl.service.AdminService.CreateNodeEvent',
    index=28,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_event__pb2._NODEEXECUTIONEVENTREQUEST,
    output_type=flyteidl_dot_admin_dot_event__pb2._NODEEXECUTIONEVENTRESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\031\"\024/api/v1/events/nodes:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='CreateTaskEvent',
    full_name='flyteidl.service.AdminService.CreateTaskEvent',
    index=29,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_event__pb2._TASKEXECUTIONEVENTREQUEST,
    output_type=flyteidl_dot_admin_dot_event__pb2._TASKEXECUTIONEVENTRESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\031\"\024/api/v1/events/tasks:\001*')),
  ),
  _descriptor.MethodDescriptor(
    name='GetTaskExecution',
    full_name='flyteidl.service.AdminService.GetTaskExecution',
    index=30,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONGETREQUEST,
    output_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTION,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\234\002\022\231\002/api/v1/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}')),
  ),
  _descriptor.MethodDescriptor(
    name='ListTaskExecutions',
    full_name='flyteidl.service.AdminService.ListTaskExecutions',
    index=31,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONLISTREQUEST,
    output_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONLIST,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\255\001\022\252\001/api/v1/task_executions/{node_execution_id.execution_id.project}/{node_execution_id.execution_id.domain}/{node_execution_id.execution_id.name}/{node_execution_id.node_id}')),
  ),
  _descriptor.MethodDescriptor(
    name='GetTaskExecutionData',
    full_name='flyteidl.service.AdminService.GetTaskExecutionData',
    index=32,
    containing_service=None,
    input_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONGETDATAREQUEST,
    output_type=flyteidl_dot_admin_dot_task__execution__pb2._TASKEXECUTIONGETDATARESPONSE,
    options=_descriptor._ParseOptions(descriptor_pb2.MethodOptions(), _b('\202\323\344\223\002\241\002\022\236\002/api/v1/data/task_executions/{id.node_execution_id.execution_id.project}/{id.node_execution_id.execution_id.domain}/{id.node_execution_id.execution_id.name}/{id.node_execution_id.node_id}/{id.task_id.project}/{id.task_id.domain}/{id.task_id.name}/{id.task_id.version}/{id.retry_attempt}')),
  ),
])
_sym_db.RegisterServiceDescriptor(_ADMINSERVICE)

DESCRIPTOR.services_by_name['AdminService'] = _ADMINSERVICE

# @@protoc_insertion_point(module_scope)