<script setup>
import { onMounted, ref } from 'vue'
import { fetchRequestInfos } from '../api/requestinfo'

let dayrecords = ref([]);
const value1 = ref(true)
const getDayrecords = async () => {
  let result = await fetchRequestInfos()
  dayrecords.value = result.data
}

onMounted(() => getDayrecords())

</script>

<template>
  <main>
    <h2>异常请求记录</h2>
    <el-table :data="dayrecords">
      <el-table-column fixed prop="url" label="访问地址" />
      <el-table-column prop="searchKey" label="关键词" />
      <el-table-column prop="requestType" label="请求类型" />
      <el-table-column prop="createdTime" label="访问时间" />
    </el-table>
  </main>
</template>
