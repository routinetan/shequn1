<template>
    <Row>
      <Row class="action"  :gutter="70">
        <Col span="1" >
          <Button type="primary" size="large" to="/group/create" >新增</Button>
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
          title: '人数',
          key: 'num',
        },
        {
          title: '展示状态',
          key: 'is_show',
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
          width: 220,
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
              }, '查看活码'),
              h('Divider', {props :{type:"vertical"}}),
              h('a', {
                on: {
                  click: () => {
                    let id = params.row.id
                    this.$router.push({path:'/group/'+id,query:params.row})
                  }
                }
              }, '编辑'),
              h('Divider', {props :{type:"vertical"}}),
              h('a', {
                on: {
                  click: () => {
                    this.remove(params.index)
                  }
                }
              }, '删除')
            ]);
          }
        }
      ],
      list: [
      ]
    }
  },
  mounted() {
    this.Request({
      url:'/groups',
      method:"GET",

    }).then(response => {
      let body =  response.data
      this.list = response.data.data
      this.page = {total:body.total,page_size:body.rows}
    })
  },
  computed: {
    iconSize () {
      return this.spanLeft === 5 ? 14 : 24;
    }
  },
  methods: {
    toggleClick () {
      if (this.spanLeft === 5) {
        this.spanLeft = 2;
        this.spanRight = 22;
      } else {
        this.spanLeft = 5;
        this.spanRight = 19;
      }
    },
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
        url:'/groups',
        method:"GET",
        params:{"num":page-1}
      }).then(response => {
        let body =  response.data
        this.list = response.data.data
        this.page = {total:body.total,page_size:body.page_size}
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
.action {
  padding: 10px 0 10px 15px;
}
.table_footer {
  padding-top:10px;
  padding-left: 10px;
}
</style>
