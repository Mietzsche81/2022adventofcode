import functools
from itertools import pairwise
from queue import PriorityQueue, Empty as EmptyException

Position = tuple[int, int]
Velocity = tuple[int, int]
Time = int

BlizzardState = tuple[Position, Velocity]
MapState = list[BlizzardState]
MapCache = dict[Time, MapState]

ObjectiveState = tuple[Time, Position]


class DepthFirstQueue(PriorityQueue):
    def __init__(self, target: Position):
        self.target: Position = target
        self.fastest_path: Time = None
        super().__init__()

    @functools.cache
    def put(self, s: ObjectiveState):
        if self.fastest_path:
            pass
        elif s[1] == self.target:
            self.fastest_path = s[0]
            self.purge_invalid()
        else:
            super().put(
                (
                    s[0] + self.distance_to_target(s[1]),  # Best score possible
                    s,  # State (t,(x,y))
                )
            )

    def get(self) -> ObjectiveState:
        return super().get()[1]

    def distance_to_target(self, pos: Position) -> int:
        return sum(
            (
                abs(self.target[0] - pos[0]),
                abs(self.target[1] - pos[1]),
            )
        )

    def purge_invalid(self):
        if self.fastest_path is None:
            return
        # Slow, but safe due to PriorityQueue infrastructure
        q: list[ObjectiveState] = []
        while not self.empty():
            try:
                t, x = self.get()
                if t < self.fastest_path:
                    q.append((t, x))
            except EmptyException:
                continue
        for s in q:
            self.put(s)


class StateIterator:
    @staticmethod
    def decode(encoding: str) -> Velocity:
        match encoding:
            case "^":
                return (-1, 0)
            case "v":
                return (1, 0)
            case ">":
                return (0, 1)
            case "<":
                return (0, -1)

    def __init__(self, grid: list[str]):
        self.X, self.Y = len(grid), len(grid[0])
        self.start: Position = (
            0,
            next(i for i, x in enumerate(grid[0]) if x == "."),
        )
        self.target: Position = (
            self.X - 1,
            next(i for i, x in enumerate(grid[-1]) if x == "."),
        )
        self._map_cache: MapCache = {
            0: [
                ((i, j), v)
                for i, line in enumerate(grid)
                for j, char in enumerate(line)
                if (v := self.decode(char))
            ]
        }
        self.dfs = DepthFirstQueue(self.target)
        self.dfs.put((0, self.start))

    @functools.cache
    def iterate_map(self, t: Time) -> MapState:
        for u in range(max(self._map_cache), t + 1):
            future: MapState = [
                (
                    (x[0] + v[0], x[1] + v[1]),
                    v,
                )
                for x, v in self._map_cache[u]
            ]
            for i, ((x, y), v) in enumerate(future):
                if x < 1:
                    future[i] = ((self.X - 2, y), v)
                elif x > (self.X - 2):
                    future[i] = ((1, y), v)
                if y < 1:
                    future[i] = ((x, self.Y - 2), v)
                elif y > (self.Y - 2):
                    future[i] = ((x, 1), v)
            self._map_cache[u + 1] = future
        return future

    @functools.cache
    def iterate_objective(self, state: ObjectiveState) -> list[ObjectiveState]:
        return [
            (t, (x, y))
            for t, (x, y) in (
                (
                    state[0] + 1,
                    (
                        state[1][0] + dx,
                        state[1][1] + dy,
                    ),
                )
                for dx, dy in [
                    (0, 0),
                    (1, 0),
                    (-1, 0),
                    (0, 1),
                    (0, -1),
                ]
            )
            if (
                ((0 < x < (self.X - 1)) and (0 < y < (self.Y - 1)))
                or (x, y) in (self.start, self.target)
            )
        ]

    def solve(self):
        while not self.dfs.empty():
            t, x = self.dfs.get()

            # Solve for/look up next environment state
            blizzards = [x for x, _ in self.iterate_map(t)]

            # Iterate objective state forward
            states = [
                (
                    u,
                    y,
                )
                for u, y in self.iterate_objective((t, x))
                if y not in blizzards
            ]
            for s in states:
                self.dfs.put(s)

        return self.dfs.fastest_path

    def solve_sequence(self, targets: list[Position]):
        t = 0
        for start, finish in pairwise(targets):
            self.dfs = DepthFirstQueue(finish)
            self.dfs.put((t, start))
            t = self.solve()

        return t
