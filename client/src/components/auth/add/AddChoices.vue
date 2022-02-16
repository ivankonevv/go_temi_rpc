<template>
  <div class="main-console-content-container">
    <div class="main-add-choice-container">
      <span class="add-choice-container-title">Добавить</span>
      <div class="add-choice-shortcut-container">
        <router-link to="/console/add/wood" class="add-choice-shortcut">
          <span class="add-choice-shortcut-title">Межкомнатная дверь</span>
        </router-link>
        <router-link to="/console/add/metal" class="add-choice-shortcut">
          <span class="add-choice-shortcut-title">Металлическая дверь</span>
        </router-link>
        <router-link to="/console/add/inventory" class="add-choice-shortcut">
          <span class="add-choice-shortcut-title">Переучет</span>
        </router-link>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "AddChoices"
}
</script>

<style scoped>
  @import "../../../assets/css/auth/console_add_chioces.css";
</style>