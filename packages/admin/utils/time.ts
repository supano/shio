import moment from 'moment'

const timeUtils = {
  timestampToDateString(t: number): string {
    const d = new Date(t)

    return moment(t).format('DD/MM/YYYY HH:mm:ss')
  }
}

export default timeUtils
