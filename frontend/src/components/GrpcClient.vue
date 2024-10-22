<template>
  <div>
      <h1>gRPC-Web Client</h1>
      <button @click="getInvoice">Get Invoice</button>
      <pre>{{ formattedResponse }}</pre>  <!-- Use formatted response for better readability -->
  </div>
</template>

<script>
import { InvoicerClient } from "../invoicer_grpc_web_pb"; // Path to generated gRPC-Web client code
import { InvoiceRequest } from "../invoicer_pb";          // Path to generated proto code

export default {
  data() {
      return {
          response: "",  // Store raw response from gRPC server
      };
  },
  computed: {
      formattedResponse() {
          // Format JSON string for better readability in the UI
          try {
              return JSON.stringify(this.response, null, 2);  // Pretty-print JSON
          } catch (e) {
              return this.response;  // Return raw response if not valid JSON
          }
      },
  },
  methods: {
      getInvoice() {
          const client = new InvoicerClient("http://localhost:8080");  // Point to Envoy proxy

          const request = new InvoiceRequest();
          request.setId("12345");  // Example invoice ID

          client.getInvoice(request, {}, (err, response) => {
              if (err) {
                  console.error("gRPC Error:", err.message);  // Log the full error message
                  this.response = `Error: ${err.message}`;
              } else {
                  console.log("gRPC Response:", response);  // Log the full response for debugging
                  const apps = response.getAppsList();  // Retrieve the list of apps
                  this.response = apps.map(app => ({
                      idApp: app.getIdApp(),
                      nameApp: app.getNameApp(),
                      descriptionApp: app.getDescriptionApp(),
                      appUrl: app.getAppUrl(),
                      imageAppUrl: app.getImageAppUrl(),
                  }));
              }
          });
      },
  },
};
</script>

<style scoped>
button {
  margin: 10px;
}
pre {
  background-color: #f5f5f5;
  padding: 10px;
  border-radius: 5px;
}
</style>
