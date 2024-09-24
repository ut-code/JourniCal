import { useCallback, useEffect, useRef, useState } from "react";
import JournalEntry from "../components/JournalEntry";
import useJournal from "../hooks/useJournal";
import { add } from "date-fns";
import TopBar from "../components/TopBar";

function Journal() {
  const {
    journals,
    isLoading,
    error,
    fetchMoreEntriesAfter,
    fetchMoreEntriesBefore,
  } = useJournal();
  const [baseDate, setBaseDate] = useState(new Date());
  const [topDate, setTopDate] = useState<Date>(
    new Date(new Date("2024-09-17").toDateString()),
  );
  const [bottomDate, setBottomDate] = useState<Date>(
    new Date(new Date("2024-09-17").toDateString()),
  );
  const topTargetRef = useRef<HTMLDivElement>(null);
  const bottomTargetRef = useRef<HTMLDivElement>(null);

  const topScrollObserver = useCallback(
    () =>
      new IntersectionObserver(
        async (entries) => {
          if (entries[0].isIntersecting) {
            await fetchMoreEntriesBefore(topDate);
            setTopDate((prev) => add(prev, { days: -4 }));
          }
        },
        {
          root: null,
          rootMargin: "200px",
          threshold: 0.01,
        },
      ),
    [fetchMoreEntriesBefore, topDate],
  );

  const bottomScrollObserver = useCallback(
    () =>
      new IntersectionObserver(
        async (entries) => {
          if (entries[0].isIntersecting) {
            await fetchMoreEntriesAfter(bottomDate);
            setBottomDate((prev) => add(prev, { days: 4 }));
          }
        },
        {
          root: null,
          rootMargin: "200px",
          threshold: 0.01,
        },
      ),
    [fetchMoreEntriesAfter, bottomDate],
  );

  useEffect(() => {
    const topTarget = topTargetRef.current;
    if (topTarget) {
      const topObserver = topScrollObserver();
      topObserver.observe(topTarget);
      return () => {
        topObserver.unobserve(topTarget);
      };
    }
  }, [topScrollObserver, topTargetRef]);

  useEffect(() => {
    const bottomTarget = bottomTargetRef.current;
    if (bottomTarget) {
      const bottomObserver = bottomScrollObserver();
      bottomObserver.observe(bottomTarget);
      return () => {
        bottomObserver.unobserve(bottomTarget);
      };
    }
  }, [bottomScrollObserver, bottomTargetRef]);

  return (
    <div className="journal-app">
      <TopBar baseDate={baseDate} setBaseDate={setBaseDate} />
      <div ref={topTargetRef} style={{ marginTop: "20%" }} />
      {isLoading ? (
        <div>Loading...</div>
      ) : error ? (
        <div>Error: {error.message}</div>
      ) : (
        <div className="journal-entries">
          {journals?.map((journal, index) => (
            <JournalEntry key={index} journal={journal} />
          ))}
        </div>
      )}
      <div ref={bottomTargetRef} />
    </div>
  );
}

export default Journal;
