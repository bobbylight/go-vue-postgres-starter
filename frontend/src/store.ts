import Vuex, { ActionContext, Store } from "vuex";
import debounce from "debounce";
import { AppState, DebounceFunction, Task, TaskReorderParams } from '@/app.d';
import Vue from "vue";
import restApi from "@/rest-api";

Vue.use(Vuex);

const taskDebounceFunctionsById: { [id: string]: DebounceFunction } = {};

const store: Store<AppState> = new Store({
    state: {
        tasks: [],
    },
    mutations: {
        addTask(state: AppState, task: Task) {
            state.tasks.push(task);
        },
        setTasks(state: AppState, tasks: Task[]) {
            state.tasks = tasks || [];
        },
        updateTask(state: AppState, task: Task) {
            const index: number = store.state.tasks.findIndex((task2: Task) => {
                return task2.id === task.id;
            });
            if (index === -1) {
                alert(`No such task: ${task.id}`);
                return;
            }
            // Construct a new array for the UI to update
            state.tasks = state.tasks
                .slice(0, index)
                .concat([task])
                .concat(state.tasks.slice(index + 1));
        },
    },
    getters: {},
    actions: {
        deleteTask(handler: ActionContext<AppState, AppState>, task: Task) {
            // Just verify that we know this task exists
            const task2: Task | undefined = store.state.tasks.find(
                (t2: Task) => {
                    return task.id === t2.id;
                }
            );
            if (!task2) {
                alert(`No such task: ${task.id}`);
                return;
            }

            restApi.deleteTask(task).then((newRemainingTasks: Task[]) => {
                store.commit("setTasks", newRemainingTasks);
            });
        },
        reorderTask(handler: ActionContext<AppState, AppState>, params: TaskReorderParams) {
            // Just verify that we know this task exists
            const task2: Task | undefined = store.state.tasks.find(
                (t2: Task) => {
                    return params.task.id === t2.id;
                }
            );
            if (!task2) {
                alert(`No such task: ${params.task.id}`);
                return;
            }

            restApi.reorderTask(params.task, params.forward).then((newRemainingTasks: Task[]) => {
                store.commit("setTasks", newRemainingTasks);
            });
        },
        updateTask(handler: ActionContext<AppState, AppState>, task: Task) {
            // Just verify that we know this task exists
            const task2: Task | undefined = store.state.tasks.find(
                (t2: Task) => {
                    return task.id === t2.id;
                }
            );
            if (!task2) {
                alert(`No such task: ${task.id}`);
                return;
            }

            restApi.updateTask(task).then((task: Task) => {
                console.log(JSON.stringify(task));
                store.commit("updateTask", task);
            });
        },

        updateTaskState(
            handler: ActionContext<AppState, AppState>,
            task: Task
        ) {
            // Just verify that we know this task exists
            const task2: Task | undefined = store.state.tasks.find(
                (t2: Task) => {
                    return task.id === t2.id;
                }
            );
            if (!task2) {
                alert(`No such task: ${task.id}`);
                return;
            }

            let func: DebounceFunction = taskDebounceFunctionsById[task.id];
            if (func) {
                (func as DebounceFunction).clear();
            }

            func = debounce(() => {
                restApi.updateTask(task).then((task: Task) => {
                    console.log(
                        `Updating task state: ${task.id} => ${task.status}`
                    );
                    store.commit("updateTask", task);
                });
            }, 200) as DebounceFunction;
            taskDebounceFunctionsById[task.id] = func;

            func();
        },
    },
});

export default store;
