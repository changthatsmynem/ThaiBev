export const maybe = <T>(value: T | null | undefined) => ({
  map: <U>(fn: (value: T) => U) => maybe(value != null ? fn(value) : null),
  getOrElse: (defaultValue: T) => value ?? defaultValue
});