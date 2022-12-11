<template>
    <Row>
      <Row class="action"  :gutter="70">
        <Col span="1" >
          <Button type="primary" size="large" :to="torul" >新增活码</Button>
        </Col>
      </Row>
      <Row>
        <Table size="large" :columns="columns1" :data="list">
        </Table>
      </Row>
      <Row class="table_footer">
        <Col offset="16" >
          <Page :total="page.total" size="small" @on-change="changePage" show-total="true"  show-elevator="true" />
        </Col>
      </Row>
    </Row>
</template>
<script>
export default {
  data () {
    return {
      spanLeft: 5,
      spanRight: 19,
      title:"",
      group_id:"",
      tourl:"",
      page:{
        total:0,
        page_size:0,
      },
      columns1: [
        {
          title: 'id',
          key: 'id',
        },
        {
          title: '群名',
          key: 'title'
        },
        {
          title: '进群人数',
          key: 'num'
        },
        {
          title: '当前群人数',
          key: 'status'
        },
        {
          title: '状态',
          key: 'status',
          render: (h, params) => {
            let status = 'success'
            let txt = '开启'
            if (params.row.is_show == 0) {
              status ='error'
              txt = '关闭'
            }
            return h('Badge', {props: {
                status:status,
                text:txt
              },},txt)
          }
        },
        {
          title: '操作',
          key: 'action',
          width: 150,
          align: 'center',
          render: (h, params) => {
            return h('div', [
              h('a', {
                on: {
                  click: (row) => {
                    let id = params.row.id
                    this.$router.push('/group/'+id+'/qrcode')
                  }
                }
              }, '编辑'),
              h('Divider', {props :{type:"vertical"}}),
              h('a', {
                on: {
                  click: () => {
                    let id = params.row.id
                    this.$router.push({path:'/group/'+id,query:params.row})
                  }
                }
              }, '删除'),
            ]);
          }
        }
      ],
      list: [
      ]
    }
  },
  mounted() {
    this.group_id = this.$route.params.id
    this.tourl = "/group/"+this.group_id+"/qrcode/add"
    this.Request({
      url:'/qrcodes',
      method:"GET",
      params:{"group_id":this.group_id}
    }).then(response => {
      let list = response.data.data
      this.list = list
    })
  },
  computed: {
    iconSize () {
      return this.spanLeft === 5 ? 14 : 24;
    }
  },
  methods: {
    show (index) {
      this.$Modal.info({
        title: 'User Info',
        content: `Name：${this.data6[index].name}<br>Age：${this.data6[index].age}<br>Address：${this.data6[index].address}`
      })
    },
    remove (index) {
      this.data6.splice(index, 1);
    },
    changePage(page) {
      this.Request({
        url:'/qrcodes',
        method:"GET",
        params:{"group_id":this.group_id,'num':page-1}
      }).then(response => {
        let list = response.data.data
        this.list = list
      })
    }
  }
}
</script>
<style scoped>
html,body{
  height:100vh;
}
.layout-ceiling-main a{
  color: #9ba7b5;
}
.layout-hide-text .layout-text{
  display: none;
}
.ivu-col{
  transition: width .2s ease-in-out;
}
.action {
  padding: 10px 0 10px 15px;
}

.table_footer {
  padding-top:10px;
  padding-left: 10px;
}
</style>
