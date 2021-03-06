<template>
  <div class="device-info-table">
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column align="center" label="Device ID" min-width="90" width="90">
        <template slot-scope="scope">
          <span>{{ scope.row.id }}</span>
        </template>
      </el-table-column>

      <el-table-column width="160px" min-width="160px" align="center" label="Last Visible Date">
        <template slot-scope="scope">
          <span>{{ scope.row.timestamp | parseTime('{y}-{m}-{d} {h}:{i}') }}</span>
        </template>
      </el-table-column>

      <el-table-column width="120px" align="center" label="Device Name">
        <template slot-scope="scope">
          <span>{{ scope.row.device_name }}</span>
        </template>
      </el-table-column>

      <el-table-column class-name="status-col" label="Online Status" width="120" min-width="120">
        <template slot-scope="{row}">
          <el-tag :type="row.status | statusFilter">
            {{ row.status }}
          </el-tag>
        </template>
      </el-table-column>

      <el-table-column min-wdith="150px" width="150px" align="center" label="IPv4 address">
        <template slot-scope="scope">
          <span>{{ scope.row.ipv4 }}</span>
        </template>
      </el-table-column>

      <el-table-column min-wdith="130px" width="130px" align="center" label="CPU usage" />

      <el-table-column min-wdith="130px" width="130px" align="center" label="Memory usage" />

      <el-table-column min-wdith="130px" width="130px" align="center" label="Received/sec" />

      <el-table-column min-wdith="130px" width="130px" align="center" label="Sent/sec" />

      <el-table-column min-wdith="130px" width="130px" align="center" label="Temperature" />

      <el-table-column min-width="200px" label="Remark">
        <template slot-scope="{row}">
          <template v-if="row.edit">
            <el-input v-model="row.title" class="edit-input" size="small" />
            <el-button
              class="cancel-btn"
              size="small"
              icon="el-icon-refresh"
              type="warning"
              @click="cancelEdit(row)"
            >
              cancel
            </el-button>
          </template>
          <span v-else>{{ row.title }}</span>
        </template>
      </el-table-column>

    </el-table>
  </div>
</template>

<script>
import { fetchList } from '@/api/article'

export default {
  name: 'DeviceInfoTable',
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 10
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    async getList() {
      this.listLoading = true
      const { data } = await fetchList(this.listQuery)
      const items = data.items
      this.list = items.map(v => {
        this.$set(v, 'edit', false) // https://vuejs.org/v2/guide/reactivity.html
        v.originalTitle = v.title //  will be used when user click the cancel botton
        return v
      })
      this.listLoading = false
    },
    cancelEdit(row) {
      row.description = row.originalTitle
      row.edit = false
      this.$message({
        message: 'The title has been restored to the original value',
        type: 'warning'
      })
    },
    confirmEdit(row) {
      row.edit = false
      row.originalTitle = row.description
      this.$message({
        message: 'The title has been edited',
        type: 'success'
      })
    }
  }
}
</script>

<style lang="scss" scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
