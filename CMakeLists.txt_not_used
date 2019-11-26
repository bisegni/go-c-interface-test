cmake_minimum_required(VERSION 3.2)

set(CMAKE_MODULE_PATH ${CMAKE_CURRENT_LIST_DIR}/)

project(dbengine)
include(ExternalProject)

message (STATUS "Adding Boost ${BOOST_VERSION} as an external project.")
set(CMAKE_CXX_STANDARD 11)
set(BOOST_VERSION 1.71.0)

if (WIN32)
     set (BOOST_BOOTSTRAP_CMD ${BOOST_GIT_REPO_DIR}/bootstrap.bat)
     # It was bjam at one point; it is now b2
    set (BOOST_BUILD_CMD ${BOOST_GIT_REPO_DIR}/b2.exe)
elseif (UNIX)
    set (BOOST_BOOTSTRAP_CMD ./${BOOST_GIT_REPO_DIR}/bootstrap.sh)
    set (BOOST_BUILD_CMD ./${BOOST_GIT_REPO_DIR}/b2)
endif ()

ExternalProject_Add (
    boost
    PREFIX ${CMAKE_CURRENT_BINARY_DIR}/boost
    GIT_REPOSITORY https://github.com/boostorg/boost.git
    GIT_TAG boost-${BOOST_VERSION}
    GIT_PROGRESS 1
    BUILD_IN_SOURCE 1
    UPDATE_COMMAND ${BOOST_BOOTSTRAP_CMD}
    CONFIGURE_COMMAND ""
    BUILD_COMMAND ${BOOST_BUILD_CMD} toolset=clang cxxflags=-stdlib=libc++ linkflags=-stdlib=libc++ --with-system link=static install --prefix=${CMAKE_CURRENT_BINARY_DIR}/boostinstall
    INSTALL_COMMAND ""
    LOG_DOWNLOAD 1
    LOG_UPDATE 1
    LOG_CONFIGURE 1
    LOG_BUILD 1
    LOG_INSTALL 1
    )

include_directories(${CMAKE_CURRENT_BINARY_DIR}/boostinstall/include)
link_directories(${CMAKE_CURRENT_BINARY_DIR}/boostinstall/lib)

set(SOURCE_FILES
    src/dbengine.cpp
    src/SqlQuery.cpp
    src/SqlColumnValue.cpp)

ADD_LIBRARY(${PROJECT_NAME} ${SOURCE_FILES})

add_dependencies(${PROJECT_NAME} boost)