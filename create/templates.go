package create

const eventTmpl = `name: "{{ .Year}}-{{ .City}}" # The name of the event. Four digit year with the city name in lower-case, with no spaces.
year: "{{ .Year}}" # The year of the event. Make sure it is in quotes.
city: "{{ .City}}" # The displayed city name of the event. Capitalize it.
event_twitter: "{{ .Twitter}}" # Change this to the twitter handle for your event such as devopsdayschi or devopsdaysmsp
description: "Devopsdays is coming to {{ .City}}!" # Edit this to suit your preferences
ga_tracking_id: "{{ .GoogleAnalytics}}" # If you have your own Google Analytics tracking ID, enter it here. Example: "UA-74738648-1"

# All dates are in unquoted YYYY-MM-DD, like this:   variable: 2016-01-05
startdate: {{ .StartDate}} # The start date of your event. Leave blank if you don't have a venue reserved yet.
enddate: {{ .StartDate}} # The end date of your event. Leave blank if you don't have a venue reserved yet.

# Leave CFP dates blank if you don't know yet, or set all three at once.
cfp_date_start: {{ .CFPDateStart}} # start accepting talk proposals.
cfp_date_end: {{ .CFPDateEnd}} # close your call for proposals.
cfp_date_announce: {{ .CFPDateAnnounce}} # inform proposers of status

cfp_open: "{{ .CFPOpen}}"
cfp_link: "{{ .CFPLink}}" #if you have a custom link for submitting proposals, add it here. This will control the Propose menu item as well as the "Propose" button.

registration_date_start: {{ .RegistrationDateStart}} # start accepting registration. Leave blank if registration is not open yet
registration_date_end: {{ .RegistrationDateEnd}} # close registration. Leave blank if registration is not open yet.

registration_closed: "{{ .RegistrationClosed }}" # set this to true if you need to manually close registration before your registration end date
registration_link: "{{ .RegistrationLink }}" # If you have a custom registration link, enter it here. This will control the Registration menu item as well as the "Register" button.

masthead_background: "{{ .MastheadBackground }}"

# Location
#
coordinates: "{{ .Coordinates }}" # The coordinates of your city. Get Latitude and Longitude of a Point: http://itouchmap.com/latlong.html
location: "{{ .Location }}" # Defaults to city, but you can make it the venue name.
#
location_address: "{{ .LocationAddress }}" #Optional - use the street address of your venue. This will show up on the welcome page if set.

nav_elements: # List of pages you want to show up in the navigation of your page.
 # - name: propose
 # - name: location
 # - name: registration
 # - name: program
 # - name: speakers
  - name: sponsor
  - name: contact
  - name: conduct
# - name: example
#   icon: "map-o" # This is a font-awesome icon that will display on small screens. Choose at http://fontawesome.io/icons/
#   url: http://mycfp.com # The url setting is optional, and only if you want the navigation to link off-site


# These are the same people you have on the mailing list and Slack channel.
team_members: # Name is the only required field for team members.
  - name: "John Doe"
  - name: "Jane Smith"
    twitter: "devopsdays"
  - name: "Sally Fields"
    employer: "Acme Anvil Co."
    github: "devopsdays"
    facebook: "https://www.facebook.com/sally.fields"
    linkedin: "https://www.linkedin.com/in/sallyfields"
    website: "https://mattstratton.com"
    image: "sally-fields.jpg"
organizer_email: "{{ .OrganizerEmail }}" # Put your organizer email address here
proposal_email: "{{ .ProposalEmail }}" # Put your proposal email address here

# List all of your sponsors here along with what level of sponsorship they have.
# Check data/sponsors/ to use sponsors already added by others.
sponsors:
  - id: samplesponsorname
    level: gold
 #  url: http://mysponsor.com/?campaign=me # Use this if you need to over-ride a sponsor URL.
  - id: arresteddevops
    level: community

sponsors_accepted : "{{ .SponsorsAccepted }}" # Whether you want "Become a XXX Sponsor!" link

# In this section, list the level of sponsorships and the label to use.
# You may optionally include a "max" attribute to limit the number of sponsors per level. For
# unlimited sponsors, omit the max attribute or set it to 0. If you want to prevent all
# sponsorship for a specific level, it is best to remove the level.
sponsor_levels:
  - id: gold
    label: Gold
#    max: 10
  - id: silver
    label: Silver
    max: 0 # This is the same as omitting the max limit.
  - id: bronze
    label: Bronze
  - id: community
    label: Community
    `

const speakerTmpl = `+++
    Title = "{{ .Title }}"
    type = "speaker"
    {{ with .Website }}website = "{{ . }}"{{ end }}
    {{ with .Twitter }}twitter = "{{ . }}"{{ end }}
    {{ with .Facebook }}facebook = "{{ . }}"{{ end }}
    {{ with .Linkedin }}linkedin = "{{ . }}"{{ end }}
    {{ with .Github }}github = "{{ . }}"{{ end }}
    {{ with .Gitlab }}gitlab = "{{ . }}"{{ end }}
    {{ with .ImagePath }}image = "{{ . }}"{{ end }}
    +++
    {{ with .Bio }}{{.}}{{ end }}
    `
