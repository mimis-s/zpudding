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
include CMakeFiles/flatsample.dir/depend.make

# Include the progress variables for this target.
include CMakeFiles/flatsample.dir/progress.make

# Include the compile flags for this target's objects.
include CMakeFiles/flatsample.dir/flags.make

samples/monster_generated.h: flatc
samples/monster_generated.h: ../samples/monster.fbs
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --blue --bold --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_1) "flatc generation: \`samples/monster.fbs\` -> \`samples/monster_generated.h\`"
	./flatc --cpp --gen-mutable --gen-object-api --reflect-names --gen-compare -o samples /home/zhangbin/work/src/flatbuffers/samples/monster.fbs

# Object files for target flatsample
flatsample_OBJECTS =

# External object files for target flatsample
flatsample_EXTERNAL_OBJECTS =

libflatsample.a: CMakeFiles/flatsample.dir/build.make
libflatsample.a: CMakeFiles/flatsample.dir/link.txt
	@$(CMAKE_COMMAND) -E cmake_echo_color --switch=$(COLOR) --green --bold --progress-dir=/home/zhangbin/work/src/flatbuffers/build/CMakeFiles --progress-num=$(CMAKE_PROGRESS_2) "Linking CXX static library libflatsample.a"
	$(CMAKE_COMMAND) -P CMakeFiles/flatsample.dir/cmake_clean_target.cmake
	$(CMAKE_COMMAND) -E cmake_link_script CMakeFiles/flatsample.dir/link.txt --verbose=$(VERBOSE)

# Rule to build all files generated by this target.
CMakeFiles/flatsample.dir/build: libflatsample.a

.PHONY : CMakeFiles/flatsample.dir/build

CMakeFiles/flatsample.dir/requires:

.PHONY : CMakeFiles/flatsample.dir/requires

CMakeFiles/flatsample.dir/clean:
	$(CMAKE_COMMAND) -P CMakeFiles/flatsample.dir/cmake_clean.cmake
.PHONY : CMakeFiles/flatsample.dir/clean

CMakeFiles/flatsample.dir/depend: samples/monster_generated.h
	cd /home/zhangbin/work/src/flatbuffers/build && $(CMAKE_COMMAND) -E cmake_depends "Unix Makefiles" /home/zhangbin/work/src/flatbuffers /home/zhangbin/work/src/flatbuffers /home/zhangbin/work/src/flatbuffers/build /home/zhangbin/work/src/flatbuffers/build /home/zhangbin/work/src/flatbuffers/build/CMakeFiles/flatsample.dir/DependInfo.cmake --color=$(COLOR)
.PHONY : CMakeFiles/flatsample.dir/depend
