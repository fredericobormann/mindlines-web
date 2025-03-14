import scenelist from '~/assets/scenelist.json'

export default defineEventHandler(event => {
  return scenelist;
})
