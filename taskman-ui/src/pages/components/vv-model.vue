<template>
  <div>
    <Modal
      v-model="modalConfig.isShow"
      :mask-closable="false"
      @on-visible-change="visibleChange"
      :title="modalConfig.title"
    >
        <Form :label-width="100" ref="form" :model="modalConfig.values">
          <FormItem
            v-for="(item, index) in modalConfig.itemConfigs"
            :key="index"
            :label="$t(item.label)"
            :prop="item.value"
            :rules="item.required ? [{ 'required': true, message: $t(item.label) + $t('can_not_be_empty'), trigger: 'change' }] : []"
          >
            <template v-if="['text', 'password'].includes(item.type)">
              <Input
                v-model="modalConfig.values[item.value]"
                style="width:90%"
                :type="item.type"
                :placeholder="item.placeholder"
              >
              </Input>
            </template>
            <template v-if="['textarea'].includes(item.type)">
              <Input
                v-model="modalConfig.values[item.value]"
                style="width:90%"
                :type="item.type"
                :rows="item.rows"
                :placeholder="item.placeholder"
              >
              </Input>
            </template>
            <template v-if="['number'].includes(item.type)">
              <InputNumber
                :max="item.max || 1000"
                :min="item.min || 1"
                v-model="modalConfig.values[item.value]"
              ></InputNumber>
            </template>
            <template v-if="['select'].includes(item.type)">
              <Select
                v-model="modalConfig.values[item.value]"
                clearable
                filterable
                style="width:90%"
                :multiple="item.multiple"
                :placeholder="item.placeholder"
              >
                <template v-for="option in modalConfig[item.options]">
                  <Option
                    :label="item.labelKey ? option[item.labelKey] : option.label"
                    :value="item.valueKey ? option[item.valueKey] : option.value"
                    :key="item.valueKey ? option[item.valueKey] : option.value"
                  >
                  </Option>
                </template>
              </Select>
            </template>
          </FormItem>
          <!-- <ValidationProvider rules="required" name="Subject" v-slot="{ errors }">
            <FormItem label="Subject">
              <Select v-model="subject" clearable placeholder="Select Subject">
                <Option label="None" value></Option>
                <Option label="Subject 1" value="s1"></Option>
                <Option label="Subject 2" value="s2"></Option>
              </Select>
            </FormItem>
          </ValidationProvider> -->
        </Form>
      <div slot="footer">
        <Button @click="cancel">{{ $t('cancel') }}</Button>
        <Button @click="ok" type="primary">{{ $t('save') }}</Button>
      </div>
    </Modal>
  </div>
</template>

<script>
import { ValidationObserver } from 'vee-validate'
// import InputWithValidation from './test/InputWithValidation'
export default {
  name: 'EForm',
  components: {
    ValidationObserver
    // InputWithValidation
  },
  props: {
    modalConfig: {
      type: Object,
      required: true,
      default: () => {
        return {}
      }
    }
  },
  data: () => ({
    subject: [],
    password: '',
    email: ''
  }),
  mounted () {},
  methods: {
    visibleChange (isHide) {
      if (!isHide) {
        this.resetValues()
        this.$refs.form.resetFields()
      }
    },
    ok () {
      this.$refs.form.validate(valid => {
        if (valid) {
          this.$emit('saveModel', JSON.parse(JSON.stringify(this.modalConfig.values)))
        }
      })
    },
    cancel () {
      this.modalConfig.isShow = false
    },
    resetValues () {
      const keys = Object.keys(this.modalConfig.values)
      keys.forEach(k => {
        let value = this.modalConfig.values[k]
        if (typeof value === 'string') {
          this.modalConfig.values[k] = ''
        }
        if (typeof value === 'object' && Array.isArray(value)) {
          this.modalConfig.values[k] = []
        }
      })
    }
  }
}
</script>
<style>
.ivu-form-item-error-tip {
  position: initial !important;
}
.ivu-form-item {
  margin-bottom: 8px;
}
</style>
