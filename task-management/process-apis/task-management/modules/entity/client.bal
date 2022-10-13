import ballerina/http;

// ------------ Tasks
public isolated function createNewTask(TasksBody payload) returns CreatedTask {
  return {
    body: {
      id: 1,
      title: payload.title,
      groupId: payload.groupId,
      status: "open"
    }
  };
}

public isolated function updateTask(string taskId, TasksTaskidBody payload) returns Task|error {
  int id = check int:fromString(taskId);
  return {
    id,
    title: payload.title,
    groupId: 1,
    status: "open"
  };
}

public isolated function deleteTask(string taskId) returns http:Ok {
  return http:OK;
}

public isolated function changeTaskGroup(string taskId, TaskidChangegroupBody payload) returns InlineResponse200{
  return {
    groupId:<int>payload.newGroupId
  };
}

public isolated function changeTaskStatus(string taskId, TaskidChangestatusBody payload) returns InlineResponse2001{
  return {
    status: payload.status
  };
}

// ------------ Groups

public isolated function createNewGroup(GroupName groupName) returns CreatedGroup {
  return {
    body: {
      id: 1,
      name: <string>groupName.name
    }
  };
}

public isolated function updateGroup(string groupId, GroupName groupName) returns Group|error {
  int id = check int:fromString(groupId);
  return {
    id,
    name: <string> groupName.name
  };
}

// ------------ Archive

public isolated function archiveTask(ArchiveTasksBody payload) returns http:Ok {
  return http:OK;
}

public isolated function archiveGroup(ArchiveGroupsBody payload) returns http:Ok {
  return http:OK;
}