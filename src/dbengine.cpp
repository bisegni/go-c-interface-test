
#include "dbengine.h"
#include "SqlQuery.h"

#include <map>
#include <string>
#include <stdio.h>
#include <iostream>
#include <boost/thread/mutex.hpp>

#define CHECK_UUID(x) \
boost::mutex::scoped_lock scoped_lock(_map_mutex); \
if(map_submitted_query.find(std::string(x)) == map_submitted_query.end()) return -1; 

typedef std::shared_ptr<SQlQuery> SqlQueryShrdPtr;
//map uuids to sql query
static boost::mutex _map_mutex;
static std::map<std::string, SqlQueryShrdPtr> map_submitted_query; 

void ACFunction() {
	printf("ACFunction()\n");
}

void init() {}

void deinit() {}

int submitQuery(const char *query, char *uuid) {
	int err = 0;
	boost::mutex::scoped_lock scoped_lock(_map_mutex);

	SqlQueryShrdPtr query_ptr = std::make_shared<SQlQuery>();
	map_submitted_query.insert(std::make_pair(query_ptr->getID(), query_ptr));

	//copy uuid within char
	query_ptr->getID().copy(uuid, query_ptr->getID().size(), 0);
	return err;
}

int columnCount(const char *uuid, int *column_count) {
	CHECK_UUID(uuid)
	*column_count = map_submitted_query[uuid]->getSchema().size();
	return 0;
}

// int getColumnType(const char *uuid, SQlType type[]) {
// 	CHECK_UUID(uuid)
// 	unsigned int i = 0;
// 	std::vector<SQlType> vec_type = map_submitted_query[uuid]->getSchema();
// 	for (auto& _type : vec_type) { 
// 		type[i++] = _type;
// 	}
// 	return 0;
// }

// int hasRow(const char *uuid, bool **has) {
// 	CHECK_UUID(uuid)
// 	return 0;
// }