import 'whatwg-fetch';
import { writable } from 'svelte/store';

export const RequestStates = Object.freeze({
  INIT: 'INIT',
  LOADING: 'LOADING',
  ERROR: 'ERROR',
  SUCCESS: 'SUCCESS',
});

async function requestJSON(method, url, params=null) {
  try {
    const result = await fetch(`${location.origin}${url}`, {
      method,
      headers: {
        'Content-Type': 'application/json'
      },
      body: params ? JSON.stringify(params) : params,
      credentials: 'same-origin',
    });
    const body = await result.json();
    if (result.status === 200 || result.status === 201) {
      return body;
    } else {
      throw new Error(body.data.message);
    }
  } catch ({ message }) {
    throw new Error(`Fetch failed for path ${url}, ${message}`);
  }
}

export const modalStore = writable({
  open: false,
  title: '',
});

export const pathsStore = writable({
  count: 0,
  paths: new Map(),
});

export const getPathsState = writable({
  state: RequestStates.INIT,
  message: 'Init.',
});
export async function getPaths() {
  getPathsState.set({
    state: RequestStates.LOADING,
    message: 'Fetching data.',
  });
  try {
    const { data: { count, paths } } = await requestJSON('GET', '/api/paths');
    getPathsState.set({
      state: RequestStates.SUCCESS,
      message: 'Success.',
    });
    pathsStore.set({
      count,
      paths: paths.reduce((acc, path) => acc.set(path.path, path), new Map()),
    });
  } catch ({ message }) {
    getPathsState.set({
      state: RequestStates.ERROR,
      message,
    });
  }
}

export const deletePathState = writable({
  state: RequestStates.INIT,
  message: 'Init.',
});
export async function deletePath(path) {
  deletePathState.set({
    state: RequestStates.LOADING,
    message: 'Deleting data.',
  });
  try {
    await requestJSON('DELETE', `/api/paths${path}`);
    deletePathState.set({
      state: RequestStates.SUCCESS,
      message: 'Success.',
    });
    pathsStore.update((prev) => {
      prev.paths.delete(path);
      return {
        count: prev.count - 1,
        paths: prev.paths,
      };
    });
  } catch ({ message }) {
    deletePathState.set({
      state: RequestStates.ERROR,
      message,
    });
  }
}