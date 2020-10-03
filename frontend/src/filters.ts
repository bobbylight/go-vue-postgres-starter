const dateFilter: Function = (dateStr: string | Date): string => {
    if (!dateStr) {
        return "";
    }
    if (typeof dateStr === "string") {
        return new Date(dateStr).toLocaleDateString();
    }
    return dateStr.toLocaleDateString();
};

export { dateFilter };
