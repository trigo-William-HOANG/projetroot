<template>
  <div id="app">
    <h1>Welcome to gRPC-Web Client</h1>
    <GrpcClient />
  </div>
</template>

<script>
import GrpcClient from './components/GrpcClient.vue';  // Import the gRPC client component

export default {
  name: 'App',
  components: {
    GrpcClient,  // Register the gRPC client component
  },
};
</script>

<style scoped>
#app {
  text-align: center;
  margin-top: 40px;
}

h1 {
  color: #42b983;
}
</style>
