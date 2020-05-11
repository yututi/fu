<template>
  <div id="app">
    <div
      class="dd"
      :class="ddClasses"
      @dragenter="dragEnter"
      @dragleave="dragLeave"
      @dragover.prevent
      @drop.prevent="dropFile"
    >
      <button class="dd__btn fu-btn fu-btn--large" @click="openFileDialog">Browse...</button>
      <span class="dd__label">or drop a image off here.</span>
    </div>
    <app-modal v-model="showPreview" title="Preview" @click="closePreview">
      <template>
        <template v-if="!uploadedObjectName">
          <img class="preview" :src="url" alt="file preview" />
          <button class="fu-btn mt-10" @click="upload">Upload</button>
        </template>
        <template v-else>
          <input type="text" readonly :value="imageRefUrl">
        </template>
      </template>
    </app-modal>
    <input type="file" class="fu-file" style="display:none;" @change="fileSelected" />
  </div>
</template>

<script>
export default {
  name: "App",
  components: {
    AppModal: () => import("./components/Modal.vue")
  },
  data: function() {
    return {
      isDragEntered: false,
      file: null,
      url: "",
      showPreview: false,
      uploadedObjectName: ""
    };
  },
  methods: {
    dragEnter: function() {
      this.isDragEntered = true;
    },
    dragLeave: function() {
      this.isDragEntered = false;
    },
    dropFile: function() {
      if (event.dataTransfer.files.length > 0) {
        var file = event.dataTransfer.files[0];

        if (file.type.startsWith("image")) {
          this.file = file;
          this.url = window.URL.createObjectURL(file);
          this.showPreview = true;
        } else {
          alert("not a image file.");
        }
        this.isDragEntered = false;
      }
    },
    closePreview: function() {},
    openFileDialog: function() {
      this.$el.querySelector(".fu-file").click();
    },
    fileSelected: function(e) {
      if (e.target.files.length > 0) {
        var file = e.target.files[0];
        this.file = file;
        this.url = window.URL.createObjectURL(file);
        this.showPreview = true;
      }
    },
    upload: async function() {
      var data = new FormData();
      data.append("file", this.file);

      var res = await fetch("/file/upload", {
        method: "POST",
        body: data
      });
      var resBody = await res.json();
      this.uploadedObjectName = resBody.name;
    }
  },
  computed: {
    ddClasses: function() {
      return {
        "dd--is-entered": this.isDragEntered
      };
    },
    imageRefUrl: function() {
      return `${window.location.href}/media/${this.uploadedObjectName}`
    }
  }
};
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  box-sizing: border-box;
  padding: 1em;
  display: flex;
  flex-direction: column;
  align-items: center;
}
.dd {
  max-width: 680px;
  width: 100%;
  height: 200px;
  border-radius: 1em;
  background-color: aliceblue;
  border: 2px dashed gainsboro;
  padding: 2em;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;

  &--is-entered {
    border-color: lightcoral;
  }

  &__label {
    color: silver;
    font-size: 2em;
    margin-top: 0.3em;
    font-weight: 100;
  }
}
.preview {
  width: 100%;
  height: auto;
}
.fu-btn {
  background-color: cornflowerblue;
  border: none;
  padding: 0.4em 1em;
  border-radius: 0.4em;
  color: white;
  cursor: pointer;
  outline: none;
  &--large {
    font-size: 1.5em;
  }
}
.mt-10 {
  margin-top: 10px;
}
</style>
