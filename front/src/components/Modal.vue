<template>
  <div :class="modalCls" class="k-modal" @click="onclick($event)">
    <div class="k-modal__dialog k-dialog k-shadow" @click.stop>
      <div class="k-dialog__header">
        <span class="k-dialog__header-title">{{title}}</span>
        <fa-icon class="k-icon" icon="times-circle" @click="onclick($event)" />
      </div>
      <div class="k-dialog__body">
        <slot />
      </div>
    </div>
  </div>
</template>
<script>
export default {
  name: "Modal",
  props: ["title", "value", "width"],
  computed: {
    show: {
      get: function() {
        return this.value;
      },
      set: function(value) {
        this.$emit("input", value);
      }
    },
    modalCls: function() {
      return {
        "k-modal--show": this.value
      };
    }
  },
  methods: {
    onclick(e) {
      this.show = false;
      this.$emit("click", e);
    }
  }
};
</script>
<style lang="scss">
@import "../styles/base.scss";
.k-modal {
  position: fixed;
  display: flex;
  top: 0px;
  left: 0px;
  align-items: center;
  justify-content: center;
  height: 100vh;
  width: 100%;
  z-index: 100;
  background-color: rgba(0, 0, 0, 0.2);
  visibility: hidden;
  opacity: 0;
  transition: 0.3s opacity;
  &--show {
    visibility: visible;
    opacity: 1;

    .k-dialog {
      transform: scale(1);
    }
  }
}

.k-dialog {
  box-sizing: border-box;
  min-width: 300px;
  width: 70%;
  max-width: 400px;
  max-height: 100%;

  @include sp {
    width: calc(100% - 10px);
    max-width: calc(100% - 10px);
  }
  transform: scale(0.2);
  display: flex;
  flex-direction: column;
  background-color: white;
  border-radius: 5px;
  overflow-x: hidden;

  transition: transform 0.1s;

  &__header {
    display: flex;
    align-items: center;
    padding: 10px;
    background-color: #3367d6;
    color: white;
  }
  &__header-title {
    flex: 1;
  }

  &__body {
    flex: 1;
    padding: 10px;
    overflow-y: auto;
  }
}
.k-icon {
  &:hover {
    cursor: pointer;
  }
}
</style>