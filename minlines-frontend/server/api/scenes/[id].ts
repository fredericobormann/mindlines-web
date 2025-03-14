import scenes from '~/assets/scenelist.json';
import scene6 from '~/assets/scenefiles/scene6.json';

export default defineEventHandler(async event => {
  const id = parseInt(getRouterParam(event, 'id') ?? 'NaN');
  if(!Number.isFinite(id)) {
    throw createError({
      status: 400,
      message: 'Scene ID must be a number',
    });
  }

  const scene = scenes.find(s => s.index === id);
  if (!scene) {
    throw createError({
      status: 404,
      message: 'Scene does not exist',
    });
  }

  return {
    ...scene,
    content: scene6,
  };
});
