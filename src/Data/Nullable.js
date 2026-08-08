/* eslint-disable no-eq-null, eqeqeq */

const nullImpl = null;
export { nullImpl as null };


export function notNull(x) {
  return x;
}
