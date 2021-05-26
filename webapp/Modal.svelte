<script>
  import {
    Button,
    Form,
    FormGroup,
    Input,
    Label,
    Modal,
    ModalBody,
    ModalFooter,
    ModalHeader,
  } from "sveltestrap";
  import { modalStore } from "./store";

  const toggle = () =>
    modalStore.update((prev) => ({
      ...prev,
      open: !prev.open,
    }));
</script>

<Modal isOpen={$modalStore.open} {toggle}>
  <ModalHeader {toggle}>{$modalStore.title}</ModalHeader>
  <ModalBody>
    <Label for="path" class="fw-bold">Path</Label>
    <Input
      type="url"
      name="path"
      id="path"
      value={$modalStore.path}
      placeholder="/example/path"
      disabled={!!$modalStore.path}
      required
    />
    <hr class="invisible" />
    <Label for="editor" class="fw-bold">Data</Label>
  </ModalBody>
  <ModalFooter>
    <Button outline color="success" on:click={toggle}>Save</Button>
    <Button outline color="secondary" on:click={toggle}>Cancel</Button>
  </ModalFooter>
</Modal>

<style>
</style>