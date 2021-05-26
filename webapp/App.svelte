<script>
  import { onMount } from 'svelte';
  import {
    Button,
    Col,
    Container,
    Icon,
    Nav,
    Navbar,
    NavbarBrand,
    NavItem,
    NavLink,
    Row,
    Spinner,
    Styles,
  } from 'sveltestrap';
  import {
    pathsStore,
    getPaths,
    getPathsState,
    deletePath,
    deletePathState,
    modalStore,
    RequestStates,
  } from './store';
  import Editor from './Editor.svelte';
  import Modal from './Modal.svelte';

  function handleDelete(path) {
    deletePath(path);
  }

  function openModal(type, path = '') {
    if (type === 'add') {
      modalStore.set({
        open: true,
        title: 'Add Path',
      });
    } else if (type === 'edit') {
      modalStore.set({
        open: true,
        title: 'Edit Path',
        path,
      });
    }
  }

  onMount(getPaths);
</script>

<Styles />

<Container sm>
  <Navbar color="light" light expand="md" class="mt-3 mb-3 row">
    <NavbarBrand href="/" class="fs-2">Hubref</NavbarBrand>
    <Nav class="ms-auto" navbar>
      <NavItem class="me-3">
        <NavLink disabled href="#">#{$pathsStore.count}</NavLink>
      </NavItem>
      <NavItem>
        <Button on:click={() => openModal("add")} color="primary">
          <Icon name="file-earmark-plus" />
        </Button>
      </NavItem>
    </Nav>
  </Navbar>
  {#if $getPathsState.state === RequestStates.SUCCESS}
    {#each [...$pathsStore.paths] as [path], i}
      <Row class="pt-3 pb-3 border-bottom border-light">
        <Col xs="auto">
          <div class="fw-bold font-monospace lh-lg pt-1">#{i + 1}</div>
        </Col>
        <Col class="pt-1">
          <a class="font-monospace lh-lg" target="_blank" href="/data{path}">{path}</a>
        </Col>
        <Col xs="auto">
          <Button on:click={() => openModal("edit", path)} color="light">
            <Icon name="file-earmark-diff" />
          </Button>
        </Col>
        <Col xs="auto">
          <Button on:click={() => handleDelete(path)} color="danger">
            <Icon name="file-earmark-x" />
          </Button>
        </Col>
      </Row>
    {/each}
  {:else if $getPathsState.state === RequestStates.LOADING}
    <Row class="pt-3 justify-content-md-center">
      <Col xs="auto">
        <Spinner class="m-auto" size="sm" type="grow" />
      </Col>
    </Row>
  {:else if $getPathsState.state === RequestStates.ERROR}
    <Row class="pt-3 justify-content-md-center">
      <Col xs="auto" class="text-danger">Error: {$getPathsState.message}</Col>
    </Row>
  {:else}
    <Row class="pt-3 justify-content-md-center">
      <Col xs="auto">Initializing...</Col>
    </Row>
  {/if}
</Container>

<Modal />

<style>
  a {
    text-decoration: none;
    color: #000000;
  }
  a:hover {
    color: #27ae60;
  }
</style>
