import defaultSettings from '@/settings'

const title = defaultSettings.title || 'Equipment Management System'

export default function getPageTitle(pageTitle) {
  if (pageTitle) {
    return `${pageTitle} - ${title}`
  }
  return `${title}`
}
