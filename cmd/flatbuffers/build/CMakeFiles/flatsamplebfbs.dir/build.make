# CMAKE generated file: DO NOT EDIT!
# Generated by "Unix Makefiles" Generator, CMake Version 3.10

# Delete rule output on recipe failure.
.DELETE_ON_ERROR:


#=============================================================================
# Special targets provided by cmake.

# Disable implicit rules so canonical targets will work.
.SUFFIXES:


# Remove some rules from gmake that .SUFFIXES does not remove.
SUFFIXES =

.SUFFIXES: .hpux_make_needs_suffix_list


# Suppress display of executed commands.
$(VERBOSE).SILENT:


# A target that is always out of date.
cmake_force:

.PHONY : cmake_force

#=============================================================================
# Set environment variables for the build.

# The shell in which to execute make rules.
SHELL = /bin/sh

# The CMake executable.
CMAKE_COMMAND = /usr/bin/cmake

# The command to remove a file.
RM = /usr/bin/cmake -E remove -f

# Escaping for special characters.
EQUALS = =

# The top-level source directory on which CMake was run.
CMAKE_SOURCE_DIR = /home/zhangbin/work/src/flatbuffers

# The top-level build directory on which CMake was run.
CMAKE_BINARY_DIR = /home/zhangbin/work/src/flatbuffers/build

# Include any dependencies generated for this target.
include CMakeFiles/flatsamplebfbs.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/flatsamplebfbs.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/flatsamplebfbs.dir/flags.make

CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o: CMakeFiles/flatsamplebfbs.dir/flags.make
CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o: ../src/idl_parser.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "Building CXX object CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o"
	/usr/bin/g++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o -c /home/zhangbin/work/src/flatbuffers/src/idl_parser.cpp

CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.i"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/zhangbin/work/src/flatbuffers/src/idl_parser.cpp > CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.i

CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.s"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/zhangbin/work/src/flatbuffers/src/idl_parser.cpp -o CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.s

CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.requires:

.PHONY : CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.requires

CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.provides: CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.requires
	$(MAKE) -f CMakeFiles/flatsamplebfbs.dir/build.make CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.provides.build
.PHONY : CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.provides

CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.provides.build: CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o


CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o: CMakeFiles/flatsamplebfbs.dir/flags.make
CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o: ../src/idl_gen_text.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Building CXX object CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o"
	/usr/bin/g++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o -c /home/zhangbin/work/src/flatbuffers/src/idl_gen_text.cpp

CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.i"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/zhangbin/work/src/flatbuffers/src/idl_gen_text.cpp > CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.i

CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.s"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/zhangbin/work/src/flatbuffers/src/idl_gen_text.cpp -o CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.s

CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.requires:

.PHONY : CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.requires

CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.provides: CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.requires
	$(MAKE) -f CMakeFiles/flatsamplebfbs.dir/build.make CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.provides.build
.PHONY : CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.provides

CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.provides.build: CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o


CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o: CMakeFiles/flatsamplebfbs.dir/flags.make
CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o: ../src/reflection.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_3) "Building CXX object CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o"
	/usr/bin/g++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o -c /home/zhangbin/work/src/flatbuffers/src/reflection.cpp

CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.i"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/zhangbin/work/src/flatbuffers/src/reflection.cpp > CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.i

CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.s"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/zhangbin/work/src/flatbuffers/src/reflection.cpp -o CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.s

CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.requires:

.PHONY : CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.requires

CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.provides: CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.requires
	$(MAKE) -f CMakeFiles/flatsamplebfbs.dir/build.make CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.provides.build
.PHONY : CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.provides

CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.provides.build: CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o


CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o: CMakeFiles/flatsamplebfbs.dir/flags.make
CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o: ../src/util.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_4) "Building CXX object CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o"
	/usr/bin/g++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o -c /home/zhangbin/work/src/flatbuffers/src/util.cpp

CMakeFiles/flatsamplebfbs.dir/src/util.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/flatsamplebfbs.dir/src/util.cpp.i"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/zhangbin/work/src/flatbuffers/src/util.cpp > CMakeFiles/flatsamplebfbs.dir/src/util.cpp.i

CMakeFiles/flatsamplebfbs.dir/src/util.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/flatsamplebfbs.dir/src/util.cpp.s"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/zhangbin/work/src/flatbuffers/src/util.cpp -o CMakeFiles/flatsamplebfbs.dir/src/util.cpp.s

CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.requires:

.PHONY : CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.requires

CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.provides: CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.requires
	$(MAKE) -f CMakeFiles/flatsamplebfbs.dir/build.make CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.provides.build
.PHONY : CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.provides

CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.provides.build: CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o


CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o: CMakeFiles/flatsamplebfbs.dir/flags.make
CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o: ../samples/sample_bfbs.cpp
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_5) "Building CXX object CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o"
	/usr/bin/g++  $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -o CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o -c /home/zhangbin/work/src/flatbuffers/samples/sample_bfbs.cpp

CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.i: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Preprocessing CXX source to CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.i"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -E /home/zhangbin/work/src/flatbuffers/samples/sample_bfbs.cpp > CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.i

CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.s: cmake_force
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green "Compiling CXX source to assembly CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.s"
	/usr/bin/g++ $(CXX_DEFINES) $(CXX_INCLUDES) $(CXX_FLAGS) -S /home/zhangbin/work/src/flatbuffers/samples/sample_bfbs.cpp -o CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.s

CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.requires:

.PHONY : CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.requires

CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.provides: CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.requires
	$(MAKE) -f CMakeFiles/flatsamplebfbs.dir/build.make CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.provides.build
.PHONY : CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.provides

CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.provides.build: CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o


# Object files for target flatsamplebfbs
flatsamplebfbs_OBJECTS = \
"CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o" \
"CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o" \
"CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o" \
"CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o" \
"CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o"

# External object files for target flatsamplebfbs
flatsamplebfbs_EXTERNAL_OBJECTS =

flatsamplebfbs: CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o
flatsamplebfbs: CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o
flatsamplebfbs: CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o
flatsamplebfbs: CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o
flatsamplebfbs: CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o
flatsamplebfbs: CMakeFiles/flatsamplebfbs.dir/build.make
flatsamplebfbs: libflatsample.a
flatsamplebfbs: CMakeFiles/flatsamplebfbs.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_6) "Linking CXX executable flatsamplebfbs"
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/flatsamplebfbs.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/flatsamplebfbs.dir/build: flatsamplebfbs

.PHONY : CMakeFiles/flatsamplebfbs.dir/build

CMakeFiles/flatsamplebfbs.dir/requires: CMakeFiles/flatsamplebfbs.dir/src/idl_parser.cpp.o.requires
CMakeFiles/flatsamplebfbs.dir/requires: CMakeFiles/flatsamplebfbs.dir/src/idl_gen_text.cpp.o.requires
CMakeFiles/flatsamplebfbs.dir/requires: CMakeFiles/flatsamplebfbs.dir/src/reflection.cpp.o.requires
CMakeFiles/flatsamplebfbs.dir/requires: CMakeFiles/flatsamplebfbs.dir/src/util.cpp.o.requires
CMakeFiles/flatsamplebfbs.dir/requires: CMakeFiles/flatsamplebfbs.dir/samples/sample_bfbs.cpp.o.requires

.PHONY : CMakeFiles/flatsamplebfbs.dir/requires

CMakeFiles/flatsamplebfbs.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/flatsamplebfbs.dir/cmake_clean.cmake
.PHONY : CMakeFiles/flatsamplebfbs.dir/clean

CMakeFiles/flatsamplebfbs.dir/depend:
	cd /home/zhangbin/work/src/flatbuffers/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/zhangbin/work/src/flatbuffers /home/zhangbin/work/src/flatbuffers /home/zhangbin/work/src/flatbuffers/build /home/zhangbin/work/src/flatbuffers/build /home/zhangbin/work/src/flatbuffers/build/CMakeFiles/flatsamplebfbs.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/flatsamplebfbs.dir/depend

