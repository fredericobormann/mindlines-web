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
  correctAnswers?: number,
}

export type Answer = 'correct' | 'wrong';

export function answerLine(line: Line, answer: Answer): Line {
  const correctAnswers = answer === 'correct' ? (line.correctAnswers ?? 0) + 1 : 0;
  return {
    ...line,
    correctAnswers,
  }
}

export type SceneContent = Scene & { content: Line[] }
