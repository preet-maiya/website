<script setup>
import WelcomeItem from './WelcomeItem.vue'
import DocumentationIcon from './icons/IconDocumentation.vue'
</script>

<template>
  <WelcomeItem>
    <template #icon>
      <DocumentationIcon />
    </template>
    <template #heading> Resume </template>

    Checkout my resume at
    <a href="/resume" target="_blank" rel="noopener">Resume</a>
    .
  </WelcomeItem>

</template>
