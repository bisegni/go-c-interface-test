#include "SqlQuery.h"
#include <boost/uuid/uuid.hpp>
#include <boost/uuid/uuid_io.hpp>
#include <boost/uuid/uuid_generators.hpp>
#include <boost/lexical_cast.hpp>

using namespace boost::uuids;

SQlQuery::SQlQuery(){
    random_generator rnd_gen;
    uuid _uuid = rnd_gen();
    query_id = boost::uuids::to_string(_uuid);
}

SQlQuery::~SQlQuery(){}

const std::string& SQlQuery::getID(){return query_id;}