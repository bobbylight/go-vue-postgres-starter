import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios";
import { DataPage, ErrorResponse, Task } from "./app";

export class RestApi {
    private readonly instance: AxiosInstance;

    constructor() {
        this.instance = axios.create({
            headers: {
                // Stops Spring Boot from challenging authenticated URLs with
                // "WWW-Authenticate: Basic header" (which triggers basic auth modal)
                "X-Requested-With": "XMLHttpRequest",
            },
        });
    }

    /**
     * Grabs the error response from the server, so we don't have to return an Axios-specific construct.
     *
     * @param {AxiosError} error The error received from the server.
     * @return The error response.
     */
    private static axiosErrorToErrorResponse(error: AxiosError): ErrorResponse {
        // AxiosError's data's payload is an ErrorResponse, but it is not a generic type
        // for some reason.  That's fine, we take extra care for non ErrorResponses too.

        if (error.response) {
            if (error.response.data.statusCode && error.response.data.message) {
                return error.response.data;
            }
            return {
                message: error.message,
                statusCode: error.response.status,
            };
        }

        console.error(
            `No response information in error: ${JSON.stringify(error)}`
        );
        return { message: error.message, statusCode: 0 };
    }

    createTask(label: string, desc: string): Promise<Task> {
        return axios
            .post("api/tasks", {
                label,
                desc,
            })
            .then((response: AxiosResponse<Task>) => {
                return response.data;
            })
            .catch((error: AxiosError) => {
                throw RestApi.axiosErrorToErrorResponse(error);
            });
    }

    deleteTask(task: Task): Promise<Task[]> {
        return axios
            .delete(`api/tasks/${task.id}`)
            .then((response: AxiosResponse<Task[]>) => {
                return response.data;
            })
            .catch((error: AxiosError) => {
                throw RestApi.axiosErrorToErrorResponse(error);
            });
    }

    getTasks(): Promise<DataPage<Task>> {
        return axios
            .get(`api/tasks?limit=10000`)
            .then((response: AxiosResponse<DataPage<Task>>) => {
                return response.data;
            })
            .catch((error: AxiosError) => {
                throw RestApi.axiosErrorToErrorResponse(error);
            });
    }

    getTaskById(taskId: string): Promise<Task> {
        return axios
            .get(`api/tasks/${taskId}`)
            .then((response: AxiosResponse<Task>) => {
                return response.data;
            })
            .catch((error: AxiosError) => {
                throw RestApi.axiosErrorToErrorResponse(error);
            });
    }

    /**
     * Moves a task forward or backward in the ordered list.  This
     * doesn't simply look at the task's "index" value since that is
     * a server-managed field and other open tasks might have their
     * "index" value updated as a result of this operation.
     *
     * @param task The task to update.
     * @param forward Whether to move the task forward or backward.
     */
    reorderTask(task: Task, forward: boolean): Promise<Task[]> {
        return axios
            .put(`api/tasks/order/${task.id}`,
                `{"forward": ${forward}}`)
            .then((response: AxiosResponse<Task[]>) => {
                return response.data;
            })
            .catch((error: AxiosError) => {
                throw RestApi.axiosErrorToErrorResponse(error);
            });
    }

    updateTask(task: Task): Promise<Task> {
        return axios
            .put(`api/tasks/${task.id}`, task)
            .then((response: AxiosResponse<Task>) => {
                return response.data;
            })
            .catch((error: AxiosError) => {
                throw RestApi.axiosErrorToErrorResponse(error);
            });
    }
}

export default new RestApi();
