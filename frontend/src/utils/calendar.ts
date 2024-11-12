import { Schedule } from "../types/types";

export type FetchedSchedule = {
  id: string;
  colorId?: string;
  start: {
    date?: string;
    dateTime?: string;
  };
  end: {
    date?: string;
    dateTime?: string;
  };
  summary: string;
};

const COLOR_DICT = [
  "#7986CB",
  "#33B679",
  "#8E24AA",
  "#E67C73",
  "#F6BF26",
  "#F4511E",
  "#039BE5",
  "#616161",
  "#3F51B5",
  "#0B8043",
  "#D50000",
];

export const scheduleFromFetchedData = (
  fetchedSchedule: FetchedSchedule,
): Schedule => {
  if (
    fetchedSchedule.start.dateTime == undefined &&
    fetchedSchedule.start.date != undefined &&
    fetchedSchedule.end.dateTime == undefined &&
    fetchedSchedule.end.date != undefined
  ) {
    return {
      id: fetchedSchedule.id,
      isAllDay: true,
      start: new Date(fetchedSchedule.start.date),
      end: new Date(fetchedSchedule.end.date),
      title: fetchedSchedule.summary,
      color:
        COLOR_DICT[
          Number(fetchedSchedule.colorId)
            ? Number(fetchedSchedule.colorId) - 1
            : 6
        ],
    };
  }
  if (
    fetchedSchedule.start.dateTime != undefined &&
    fetchedSchedule.start.date == undefined &&
    fetchedSchedule.end.dateTime != undefined &&
    fetchedSchedule.end.date == undefined
  ) {
    return {
      id: fetchedSchedule.id,
      isAllDay: false,
      start: new Date(fetchedSchedule.start.dateTime),
      end: new Date(fetchedSchedule.end.dateTime),
      title: fetchedSchedule.summary,
      color:
        COLOR_DICT[
          Number(fetchedSchedule.colorId)
            ? Number(fetchedSchedule.colorId) - 1
            : 6
        ],
    };
  }
  throw new Error("invalid schedule format.");
};
