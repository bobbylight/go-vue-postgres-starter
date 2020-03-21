import axios, { AxiosError, AxiosInstance, AxiosResponse } from "axios";
import { ErrorResponse, Widget } from "./app";

export class RestApi {
    private readonly instance: AxiosInstance;

    constructor() {
        this.instance = axios.create({
            headers: {
                // Stops Spring Boot from challenging authenticated URLs with
                // "WWW-Authenticate: Basic header" (which triggers basic auth modal)
                "X-Requested-With": "XMLHttpRequest"
            }
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
                statusCode: error.response.status
            };
        }

        console.error(`No response information in error: ${JSON.stringify(error)}`);
        return { message: error.message, statusCode: 0 };
    }

    getWidget(widgetId: string): Promise<Widget> {

        return axios.get(`api/widgets/${widgetId}`)
            .then((response: AxiosResponse<Widget>) => {
                return response.data;
            })
            .catch((error: AxiosError) => {
                throw RestApi.axiosErrorToErrorResponse(error);
            });
    }
}

export default new RestApi();
