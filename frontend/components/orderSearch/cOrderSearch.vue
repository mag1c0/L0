<script lang="ts" setup>
const data = ref('')
const orderId = ref('')
const error = ref('')
const resultBlock = ref(<any>undefined)

function getOrder() {
  error.value = ''
  data.value = ''
  uSmooth(resultBlock.value, -250)

  if (orderId.value === "") {
    error.value = "Неверный номер заказа"
    return
  }
  $fetch(`/api/v1/orders/` + orderId.value, {
    method: 'GET',
    baseURL: 'http://localhost:30016',
  }).then(function (res:any) {
    data.value = res;
  }).catch(()=>{
    error.value = "Ошибка при получении данных."
  })
}
</script>

<template>
  <div>
    <h2 class="heading">
      <span class="">может,</span>
      <span class="desktop offset">отследим заказ?</span>
    </h2>

    <div class="form">
      <label class="label">
        Номер заказа
        <input class="input" v-model.trim="orderId" placeholder="1140de4d-9173-4b8d-9afd-79d1b82732c6" type="text">
      </label>

      <section class="buttonSection">
        <button class="button" @click="getOrder">Отследить</button>
      </section>
    </div>

    <div ref="resultBlock" class="result">
      <p class="error" v-if="error !==''">{{ error }}</p>
      <pre :class="{show: error === '' && data}">{{ data }}</pre>
    </div>
  </div>
</template>

<style scoped>
.heading {
  position: relative;
  top: -140px;
  z-index: 3;
  display: flex;
  flex-direction: column;
  color: #fff;
  font-size: 58px;
  line-height: 54px;
  letter-spacing: -2.5px;
  text-transform: uppercase;
  transform: skew(-0, -15deg);

}

.offset {
  margin-left: 38px;
}

.label {
  width: 100%;
  font-size: 20px;
  line-height: 21px;
  letter-spacing: -.01em;
}

.form {
  padding-bottom: 30px;
}

.input {
  width: 100%;
  border: none;
  background-color: transparent;
  color: #fff;
  font-size: 48px;
  line-height: 45px;
  letter-spacing: -.01em;
  border-bottom: 1px solid hsla(0, 0%, 100%, .5);
  padding-bottom: 4px;
  text-decoration: none;
  outline: none;
}

.buttonSection {
  display: flex;
  flex-direction: column;
  gap: 5px;
  margin-top: 30px;
  margin-bottom: 15px;
}

.button {
  font-size: 120px;
  line-height: 110px;
  letter-spacing: -.053em;
  text-transform: uppercase;
  padding: 5px 7px 3px 0;
  color: #000;
  background: linear-gradient(105.06deg, #ff067e -5.12%, #be06ff 119.05%);
  border: none;
  outline: none;
}

.button:hover {
  background: #9eff00;
}

.result {
  min-height: 1049px;
}

.error {
  color: red;
  font-size: 20px;
}

pre.show {
  box-sizing: border-box;
  width: 100%;
  padding: 0;
  margin: 0;
  overflow: auto;
  overflow-y: hidden;
  font-size: 14px;
  line-height: 21px;
  background: #efefef;
  padding: 10px;
  color: #333;
}
</style>