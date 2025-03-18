export type Scene = {
  name: string,
  identifier: string,
  index: number,
};

export const characters = ['BIEDERMANN', 'BABETTE', 'SCHMITZ', 'EISENRING', 'ANNA', 'CHOR', 'CHORFÜHRER', 'DR. PHIL'] as const;

export type Character = typeof characters[number];

export type Line = {
  character: Character,
  line: string,
  reviewTimes?: {
    again: string,
    hard: string,
    good: string,
    easy: string,
  },
  dueTime: Date,
}

export const answers = {
  again: 1,
  hard: 2,
  good: 3,
  easy: 4,
};

export type Answer = keyof typeof answers;

export type SceneContent = Scene & { content: Line[] }
