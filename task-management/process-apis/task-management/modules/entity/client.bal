import ballerina/http;

// ------------ Tasks
public isolated function createNewTask(int groupId, TaskTitle taskTitle) returns CreatedTask {
  return {
    body: {
      id: 1,
      title: <string>taskTitle.title,
      groupId: groupId,
      status: "open"
    }
  };
}

public isolated function updateTask(string taskId, int groupId, TaskTitle taskTitle) returns Task|error {
  int id = check int:fromString(taskId);
  return {
    id,
    title: <string>taskTitle.title,
    groupId: groupId,
    status: "open"
  };
}

public isolated function deleteTask(string taskId) returns http:Ok {
  return http:OK;
}

public isolated function changeTaskGroup(string taskId, int groupId, int newGroupId) returns InlineResponse200{
  return {
    groupId:newGroupId
  };
}

public isolated function changeTaskStatus(int groupId, string status) returns InlineResponse2001{
  return {
    status
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

public isolated function archiveTask(int groupId, int taskId) returns http:Ok {
  return http:OK;
}

public isolated function archiveGroup(int groupId) returns http:Ok {
  return http:OK;
}