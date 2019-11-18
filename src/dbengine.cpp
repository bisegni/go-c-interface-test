
#include "dbengine.h"
#include "SqlQuery.h"

#include <map>
#include <string>
#include <stdio.h>

#include <boost/thread/mutex.hpp>


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
	if(map_submitted_query.find(query) != map_submitted_query.end()) return -1;

	SqlQueryShrdPtr query_ptr = std::make_shared<SQlQuery>();
	map_submitted_query.insert(std::make_pair(query_ptr->getID(), query_ptr));

	//copy uuid within char
	query_ptr->getID().copy(uuid, query_ptr->getID().size(), 0);
	return err;
}