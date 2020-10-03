export interface AppState {
    tasks: Task[];
}

export interface DataPage<T> {
    offset: number;
    data: T[];
    count: number;
}

export interface DebounceFunction extends Function {
    clear(): void;
}

export interface Task {
    id: string;
    label: string;
    desc: string;
    index: number;
    status: TaskStatus;
    createdAt: string;
    completedAt?: string;
}

export interface TaskReorderParams {
    task: Task;
    forward: boolean;
}

export type TaskStatus = "notStarted" | "inProgress" | "completed";

export interface ErrorResponse {
    statusCode: number;
    message: string;
}
