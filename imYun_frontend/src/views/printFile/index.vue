<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.key" :placeholder="$t('请输入文件名称或者编号')" style="width: 200px; margin-right: 10px" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-date-picker
        v-model="listQuery.dateFrom"
        type="date"
        placeholder="起始日期"
        value-format="yyyy-MM-dd HH:mm:ss"
        class="filter-item"
      />
      <el-date-picker
        v-model="listQuery.dateEnd"
        type="date"
        placeholder="截止日期"
        value-format="yyyy-MM-dd HH:mm:ss"
        class="filter-item"
      />
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" style="margin-right: 10px" @click="handleFilter">
        {{ $t('table.search') }}
      </el-button>
      <el-button class="filter-item" type="primary" icon="el-icon-download" style="margin-right: 10px" @click="downloadAll(listChecked)">
        批量下载
      </el-button>
      <el-button class="filter-item" type="danger" icon="el-icon-delete" style="margin-right: 10px" @click="handleDeleteOrders()">
        批量删除
      </el-button>
      <el-button class="filter-item" type="success" icon="el-icon-refresh-right" style="margin-right: 10px" @click="getList()">
        刷新
      </el-button>
    </div>
    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      :reserve-selection="true"
      :row-class-name="tableRowClassName"
      :default-sort="{prop: 'date', order: 'descending'}"
      @selection-change="handleSelectionChange"
    >
      <el-table-column
        type="selection"
        width="55"
      />
      <el-table-column :label="$t('ID')" prop="id" align="center" width="80" :class-name="getSortClass('id')">
        <template slot-scope="{row}">
          <span>{{ row.id }}</span>
        </template>
      </el-table-column>
      <el-table-column :label="$t('文件标题')" min-width="80px" height="10px">
        <template slot-scope="{row}">
          <el-card class="box-card" @click="handleUpdate(row)">
            <div slot="header" class="clearfix">
              <span>{{ row.fileName }}</span>
              <el-button style="float: right; padding: 3px 0" type="text">
                {{ row.fileName.match(/[^\.[^\\/]+$/)[0] }}
              </el-button>
            </div>
            <div class="text item">
              <span>
                <el-tag
                  type=""
                  size="small"
                  :hit="true"
                  effect="plain"
                >
                  {{ row.fileColor === 0 ? '黑白':'彩色' }}
                </el-tag>
              </span>
              <el-divider direction="vertical" />
              <el-tag
                type=""
                size="small"
                :hit="true"
                effect="plain"
              >
                {{ row.direction === 0 ? '纵向':'横向' }}
              </el-tag>
              <el-divider direction="vertical" />
              <el-tag
                type=""
                size="small"
                :hit="true"
                effect="plain"
              >
                {{ row.fileNum }}
              </el-tag>
              <el-divider direction="vertical" />
              <el-tag
                type=""
                size="small"
                :hit="true"
                effect="plain"
              >
                {{ row.singleSide === 0 ? '单面打印': '双面打印' }}
              </el-tag>
              <el-divider direction="vertical" />
              <el-tag
                type=""
                size="small"
                :hit="true"
                effect="plain"
              >
                {{ row.paperFormat }}
              </el-tag>
              <el-divider direction="vertical" />
              <span v-if="row.remarks !== ''">备注：{{ row.remarks }}</span>
              <el-button style="float: right; padding: 3px 0" type="text">
                <svg-icon :icon-class="fileType(row)" style="font-size: 30px" />
              </el-button>
            </div>
          </el-card>
        </template>
      </el-table-column>
      <el-table-column :label="$t('用户昵称')" width="110px" align="center">
        <template slot-scope="{row}">
          <span>{{ row.nickName }}</span>
        </template>
      </el-table-column>
      <!-- <el-table-column :label="$t('手机号')" align="center" width="110px">
        <template slot-scope="{row}">
          <span class="link-type">{{ row.userPhone }}</span>
        </template>
      </el-table-column> -->
      <el-table-column :label="$t('编号')" align="center" width="95">
        <template slot-scope="{row}">
          <span class="link-type">{{ row.code }}</span>
        </template>
      </el-table-column>

      <el-table-column :label="$t('table.actions')" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button v-if="fileType(row) === 'image'" type="success" size="mini" icon="el-icon-printer" @click="handleImage(row)">
            {{ $t('图片-打印') }}
          </el-button>
          <el-link v-if="fileType(row) !== 'image'" target="_break" :href="printUrl" :underline="false" style="margin-left:15px">
            <el-button v-if="(fileType(row) !== 'unknown' && (fileType(row) !== 'pdf'))" type="success" size="mini" icon="el-icon-printer" @click="handlePrint(row)">
              {{ $t('预览-打印') }}
            </el-button>
          </el-link>
          <el-link target="_break" :href="printUrl" :underline="false" style="margin-left:15px">
            <el-button v-if="fileType(row) === 'pdf'" type="success" size="mini" icon="el-icon-printer" @click="handlePdf(row)">
              {{ $t('预览-打印') }}
            </el-button>
          </el-link>
          <div>
            <el-link target="_self" :href="printUrl" :underline="false" style="margin-left:15px">
              <el-button type="primary" icon="el-icon-download" size="mini" @click="attDownload(row, $index)">下载</el-button>
            </el-link>
          </div>
          <el-button type="text" size="mini" @click="handleDeleteOrder(row,$index)">删除<i class="el-icon-delete" /></el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :offset.sync="listQuery.Offset" :limit.sync="listQuery.limit" @pagination="getPagination" />

    <el-dialog
      title="图片预览"
      :visible.sync="imageDialog"
      width="70%"
      center
    >
      <el-image :src="imgUrl" :preview-src-list="imgUrlList">
        <div slot="placeholder" class="image-slot">
          加载中<span class="dot">...</span>
        </div>
      </el-image>
      <span slot="footer" class="dialog-footer">
        <el-button-group>
          <el-button type="primary" @click="imageDialog = false">取 消</el-button>
          <el-button type="primary" @click="handleDeleteOrder(row,$index)">删 除</el-button>
          <el-button type="primary" @click="imageDialog = false">下 载</el-button>
          <el-button type="primary">打 印</el-button>
        </el-button-group>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import neffos from 'neffos.js'
import axios from 'axios'
import fileDownload from 'js-file-download'
import { deleteOrder, deleteOrders } from '@/api/order'
import { fetchList, updateArticle } from '@/api/file-print'
import waves from '@/directive/waves' // waves directive
// import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination
import { getToken } from '@/utils/auth'
import GLOBAL from '@/api/global_variable'

export default {
  name: 'ComplexTable',
  components: { Pagination },
  directives: { waves },
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
      wsURL: GLOBAL.wsURL,
      uploadCallbackInfo: undefined, // 接收websocket后端传回参数
      // 分页传参
      paginationList: {
        Limit: 10,
        Offset: 0,
        Key: '',
        DateFrom: '',
        DateEnd: ''
      },
      test: {},
      idIndex: [],
      idList: [],
      // 批量删除订单ID列表
      deleteOrderList: [],
      deleteOrderTemp: {}, // 批量删除订单临时存放id
      tableKey: 0,
      list: null,
      listChecked: [],
      delList: [],
      total: 0,
      listLoading: true,
      listQuery: {
        Limit: 10,
        Offset: 0,
        Key: '',
        dateFrom: '',
        dateEnd: ''
      },
      printUrl: null,
      imgUrl: 'null',
      imgUrlList: [],
      imageDialog: false,
      importanceOptions: [1, 2, 3],
      multipID: [],
      ossUrl: GLOBAL.ossUrl,
      downUrl: GLOBAL.downUrl,
      statusOptions: ['published', 'draft', 'deleted'],
      showReviewer: false,
      temp: {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        type: '',
        status: 'published'
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogPvVisible: false,
      pvData: [],
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      downloadLoading: false
    }
  },
  created() {
    this.getList()
    this.websocket()
  },
  methods: {
    // websocket
    async websocket() {
      const that = this
      const token = getToken()

      // wsURL += "?token=" + token;
      // wsURL += "?token=" + getToken()

      try {
        const conn = await neffos.dial(that.wsURL, {
          default: { // "default" namespace.
            _OnNamespaceConnected: (nsConn, msg) => {
              that.handleNamespaceConnectedConn(nsConn)
            },
            _OnNamespaceDisconnect: (nsConn, msg) => {
            },
            _OnAnyEvent: function(nsConn, msg) { // any event.
              if (msg.Event === 'Notify') {
                that.uploadCallbackInfo = msg.Body
                that.list.unshift(JSON.parse(msg.Body))
              }
            }
          }
        }, {
          headers: { 'Authorization': token }
        })
        conn.connect('default')
      } catch (err) {
        this.handleError(err)
      }
    },
    handleNamespaceConnectedConn(nsConn) {
      nsConn.emit('Hello from browser client side!')
      const token = getToken()
      nsConn.emit('Authorization', token)
    },
    // 下载
    attDownload(row, index) {
      const path = this.downUrl
      // downloadFile(fileName).then((res) => {
      //   fileDownload(res.data, row.filesUrl)
      // })
      // var path = this.ossUrl + row.rename.substring(0, 6) + '/' + row.rename
      axios({
        url: path,
        method: 'get',
        headers: {
          'Authorization': getToken()
        },
        params: {
          ReFileName: row.rename
        },
        responseType: 'blob'
      }).then((res) => {
        fileDownload(res.data, row.fileName)
      })
    },
    // 有此方法才能在handleSelectionChange方法中获取到选中行的index
    tableRowClassName({ row, rowIndex }) {
      row.index = rowIndex
    },
    // 多选table的id存储到idList
    handleSelectionChange(val) {
      this.idList = []
      this.idIndex = []
      val.forEach(item => {
        this.idList.push(item.id)
        this.idIndex.push(item.index)
      })
    },
    getList() {
      this.listLoading = true
      fetchList(this.listQuery).then(response => {
        this.listLoading = false
        this.list = response.data.items
        this.total = response.data.total
      }).catch(function(error) {
        console.log(error)
      })
    },
    // 分页传参
    getPagination(info) {
      this.listLoading = true
      this.paginationList.Limit = info.limit
      if (info.offset === 1) {
        this.paginationList.Offset = 0
      } else {
        // this.paginationList.Limit *= offset.page
        this.paginationList.Offset = (info.page - 1) * info.limit
      }
      fetchList(this.paginationList).then(response => {
        this.listLoading = false
        this.list = response.data.items
        this.total = response.data.total
      }).catch(function(error) {
        console.log(error)
      })
    },
    handleFilter() {
      this.listQuery.Offset = 0
      this.getList()
    },
    sortChange(data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID(order) {
      if (order === 'ascending') {
        this.listQuery.sort = '+id'
      } else {
        this.listQuery.sort = '-id'
      }
      this.handleFilter()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        status: 'published',
        type: ''
      }
    },
    updateData() {
      this.$refs['dataForm'].validate((valid) => {
        if (valid) {
          const tempData = Object.assign({}, this.temp)
          tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
          updateArticle(tempData).then(() => {
            const index = this.list.findIndex(v => v.id === this.temp.id)
            this.list.splice(index, 1, this.temp)
            this.dialogFormVisible = false
            this.$notify({
              title: '成功',
              message: '更新成功',
              type: 'success',
              duration: 2000
            })
          })
        }
      })
    },
    handlePrint(row) {
      const microsoftUrl = 'https://view.officeapps.live.com/op/view.aspx?src='
      this.printUrl = microsoftUrl + this.ossUrl + row.rename.substr(0, 6) + '/' + row.rename
      console.log(row.rename)
      console.log(this.printUrl)
    },
    handlePdf(row) {
      const xdoc = 'https://view.xdocin.com/xdoc?_xdoc='
      this.printUrl = xdoc + this.ossUrl + row.rename.substr(0, 6) + '/' + row.rename
      console.log(this.printUrl)
    },
    handleImage(row) {
      this.imageDialog = true
      this.imgUrl = this.ossUrl + row.rename.substr(0, 6) + '/' + row.rename
      this.imgUrlList = this.ossUrl + row.rename.substr(0, 6) + '/' + row.rename
      console.log(this.imgUrl)
    },
    downloadNotify() {
      this.$notify({
        title: '成功',
        message: '下载成功',
        type: 'success',
        duration: 2000
      })
    },
    handleDeleteOrders() {
      const obj = {}
      obj['orderId'] = this.idList
      const temp = JSON.stringify(obj)
      console.log(obj)
      console.log(temp)
      // Object.values(that.idList).forEach(value => {
      //   that.deleteOrderTemp['orderId'] = value
      //   that.deleteOrderList.push(that.deleteOrderTemp)
      //   that.deleteOrderTemp = {}
      // })
      deleteOrders(this.idList).then(() => {
        this.$notify({
          title: '成功',
          message: '删除成功',
          type: 'success',
          duration: 3000
        })
        // this.list.splice(index, 1)
        // this.imageDialog = false
      })
    },
    // 单一删除订单
    handleDeleteOrder(row, index) {
      const tmpId = this.list[index]
      console.log(tmpId)
      deleteOrder(tmpId.id).then(() => {
        this.$notify({
          title: '成功',
          message: '删除成功',
          type: 'success',
          duration: 3000
        })
        this.list.splice(index, 1)
        this.imageDialog = false
      })
    },
    downloadAll(row) {
      const _this = this
      if (_this.listChecked.length === 0) {
        this.$notify({
          title: '失败',
          message: 'Sorry,你还没有选择要下载的文件哦',
          type: 'error',
          duration: 3000
        })
        return
      }
    },
    deleteAll(row) {
      if (this.listChecked.length === 0) {
        this.$notify({
          title: '失败',
          message: 'Sorry,你还没有选择要删除的文件哦',
          type: 'error',
          duration: 3000
        })
        return
      }
      for (let i = this.list.length; i > 0; i--) {
        for (let j = 0; j < this.listChecked.length; j++) {
          if (
            this.list[i - 1] === this.listChecked[j]
          ) {
            this.list.splice(i - 1, 1)
          }
        }
      }
      // const param = {
      //   'token': X-Token('token'),
      //   'listCheckedList': this.listChecked
      // }

      console.log(this.listChecked.length)
      // this.list.splice(this.listChecked, this.listChecked.length)
      // _this.$confirm('是否确认此操作?', '提示', {
      //   confirmButtonText: '确定',
      //   cancelButtonText: '取消',
      //   type: 'warning'
      // }).then(() => {
      //   row.forEach(element => {
      //     _this.ids.push(element.chargingStationId)
      //   })
      // const param = {
      //   'token': getSessiontoken('token'),
      //   'chargingStationIdList': _this.ids
      // }
      // deleteAllCharging(param).then(function(res) {
      //   var obj = JSON.parse(utilFile.decrypt(res.data.a))
      //   if (obj.code == '200') {
      //     _this.$message.success('操作成功')
      //     _this.chargingUserList()
      //   } else {
      //     _this.$message.error(obj.msg)
      //   }
      // }).catch(function(err) {
      //   console.log(err)
      // })
      // }).catch(() => {
      //   alert(2)
      //   this.$message({
      //     type: 'info',
      //     message: '已取消'
      //   })
      // })
    },

    // formatJson(filterVal) {
    //   return this.list.map(v => filterVal.map(j => {
    //     if (j === 'timestamp') {
    //       return parseTime(v[j])
    //     } else {
    //       return v[j]
    //     }
    //   }))
    // },
    fileType(row) {
      const sourceType = row.fileName.match(/[^\.[^\\/]+$/)[0]
      const imageType = ['jpg', 'jpeg', 'png', 'bmp', 'tiff', 'svg', 'exif']
      const docType = ['docx', 'wps', 'doc']
      const execlType = ['xlsx', 'xls', 'csv', 'xlt']
      if (sourceType === 'pdf') {
        return 'pdf'
      // eslint-disable-next-line no-constant-condition
      } else if (execlType.indexOf(sourceType) !== -1) {
        return 'execl'
      } else if (sourceType === 'pptx') {
        return 'ppt'
      } else if (docType.indexOf(sourceType) !== -1) {
        return 'word'
      } else if (imageType.indexOf(sourceType) !== -1) {
        return 'image'
      } else if (sourceType === 'psd') {
        return 'psd'
      } else {
        return 'unknown'
      }
    },
    getSortClass: function(key) {
      const sort = this.listQuery.sort
      return sort === `+${key}` ? 'ascending' : 'descending'
    }
  }
}
</script>
