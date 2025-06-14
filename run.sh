#!/bin/bash

# 定义可配置变量
project_name="go-layout"
base_path="/opt/app/firefly/microservice"
node_path="${base_path}/${project_name}/main"
log_file="./nohup.log"

# 日志函数
log() {
  local message=$1
  echo "$(date '+%Y-%m-%d %H:%M:%S') - $message"
}

# 函数：精确查找进程PID
find_pid() {
  local target_path=$1
  pgrep -f -- "${target_path}"
}

# 函数：安全停止进程
safe_kill() {
  local pid=$1
  if [ -n "$pid" ]; then
    log "尝试优雅停止进程: $pid"
    kill -15 "$pid"  # 优雅停止
    sleep 5
    if ps -p "$pid" > /dev/null; then
      log "进程未停止，强制终止: $pid"
      kill -9 "$pid"
    fi
  fi
}

# 步骤1: 停止旧节点
log "停止旧节点..."
node_pid=$(find_pid "${node_path}")
safe_kill "$node_pid"
log "旧节点PID: ${node_pid}"

# 步骤2: 启动新节点
log "启动新节点..."
nohup "${node_path}" > "${log_file}" 2>&1 &
node_pid=$!

# 检查新节点是否启动成功
sleep 5
if ps -p "${node_pid}" > /dev/null; then
  log "新节点启动成功，PID: ${node_pid}"
else
  log "新节点启动失败，请检查日志: ${log_file}"
  exit 1
fi

log "服务更新完成"