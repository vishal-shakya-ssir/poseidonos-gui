#!/bin/bash
# 
# run_ibofos.sh
#
logfile="ibofos.log"
root_dir=/root/workspace/ibofos
ibofos=${root_dir}/bin/ibofos
NR_VOLUME=1
SPDK_DIR=/root/workspace/ibofos/lib/spdk
TRANSPORT=TCP
IP=10.100.11.7

pkill -9 ibofos
sleep 2

if [ ! -f "$ibofos" ]; then
        echo "fail to find $ibofos. run make prior to run run_os.sh"
else
#       ${root_dir}/script/setup_env.sh
        \rm -rf /dev/shm/ibof_nvmf_trace*
#       ${root_dir}/script/m9k/cpumode.sh max
        nohup $ibofos primary &>> ${logfile} &
        sleep 2
#	sudo ${root_dir}/lib/spdk-19.01.1/scripts/rpc.py construct_malloc_bdev -b uram0 1024 512
#	sudo ${root_dir}/lib/spdk-19.01.1/scripts/rpc.py nvmf_subsystem_create nqn.2019-04.ibof:subsystem1 -a -s IBOF00000000000001
#	sudo ${root_dir}/lib/spdk-19.01.1/scripts/rpc.py nvmf_create_transport -t RDMA -u 131072 -p 4 -c 0
#	sudo ${root_dir}/lib/spdk-19.01.1/scripts/rpc.py nvmf_subsystem_add_listener nqn.2019-04.ibof:subsystem1 -t RDMA -a 172.16.1.1 -s 1158

	sudo $SPDK_DIR/scripts/rpc.py nvmf_create_subsystem nqn.2019-04.ibof:subsystem1 -a -s IBOF00000000000001 -d IBOF_VOLUME_EXTENTION -m 512
	sudo $SPDK_DIR/scripts/rpc.py nvmf_create_subsystem nqn.2019-04.ibof:subsystem2 -a -s IBOF00000000000002 -d IBOF_VOLUME_EXTENTION_2 -m 512
	sudo $SPDK_DIR/scripts/rpc.py bdev_malloc_create -b uram0 4096 512
	sudo $SPDK_DIR/scripts/rpc.py bdev_malloc_create -b uram1 4096 512
	sudo $SPDK_DIR/scripts/rpc.py nvmf_create_transport -t $TRANSPORT -u 131072 -p 4 -c 0
	sudo $SPDK_DIR/scripts/rpc.py nvmf_subsystem_add_listener nqn.2019-04.ibof:subsystem1 -t $TRANSPORT -a $IP -s 1158
	sudo $SPDK_DIR/scripts/rpc.py nvmf_subsystem_add_listener nqn.2019-04.ibof:subsystem2 -t $TRANSPORT -a $IP -s 1158
	sudo $SPDK_DIR/scripts/rpc.py nvmf_get_subsystems
fi
