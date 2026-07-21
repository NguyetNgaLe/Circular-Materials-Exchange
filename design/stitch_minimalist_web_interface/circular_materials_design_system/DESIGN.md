---
name: Circular Materials Design System
colors:
  surface: '#faf9f9'
  surface-dim: '#dadada'
  surface-bright: '#faf9f9'
  surface-container-lowest: '#ffffff'
  surface-container-low: '#f4f3f3'
  surface-container: '#eeeeed'
  surface-container-high: '#e9e8e8'
  surface-container-highest: '#e3e2e2'
  on-surface: '#1a1c1c'
  on-surface-variant: '#42493b'
  inverse-surface: '#2f3131'
  inverse-on-surface: '#f1f0f0'
  outline: '#727a69'
  outline-variant: '#c1c9b6'
  surface-tint: '#336b07'
  primary: '#336b07'
  on-primary: '#ffffff'
  primary-container: '#70ad47'
  on-primary-container: '#193d00'
  inverse-primary: '#98d86c'
  secondary: '#2b5cad'
  on-secondary: '#ffffff'
  secondary-container: '#7ca8fe'
  on-secondary-container: '#003b82'
  tertiary: '#136299'
  on-tertiary: '#ffffff'
  tertiary-container: '#63a2dd'
  on-tertiary-container: '#00375b'
  error: '#ba1a1a'
  on-error: '#ffffff'
  error-container: '#ffdad6'
  on-error-container: '#93000a'
  primary-fixed: '#b3f585'
  primary-fixed-dim: '#98d86c'
  on-primary-fixed: '#0a2100'
  on-primary-fixed-variant: '#235100'
  secondary-fixed: '#d8e2ff'
  secondary-fixed-dim: '#adc6ff'
  on-secondary-fixed: '#001a41'
  on-secondary-fixed-variant: '#004494'
  tertiary-fixed: '#cfe5ff'
  tertiary-fixed-dim: '#98cbff'
  on-tertiary-fixed: '#001d33'
  on-tertiary-fixed-variant: '#004a77'
  background: '#faf9f9'
  on-background: '#1a1c1c'
  surface-variant: '#e3e2e2'
  leaf-green: '#70AD47'
  ocean-blue: '#4472C4'
  sky-info: '#5B9BD5'
  earth-gray: '#333333'
  surface-white: '#FFFFFF'
  warning-amber: '#FFC000'
  exchange-orange: '#ED7D31'
typography:
  display-lg:
    fontFamily: Inter
    fontSize: 48px
    fontWeight: '700'
    lineHeight: 56px
    letterSpacing: -0.02em
  headline-lg:
    fontFamily: Inter
    fontSize: 32px
    fontWeight: '600'
    lineHeight: 40px
    letterSpacing: -0.01em
  headline-lg-mobile:
    fontFamily: Inter
    fontSize: 28px
    fontWeight: '600'
    lineHeight: 36px
  headline-md:
    fontFamily: Inter
    fontSize: 24px
    fontWeight: '600'
    lineHeight: 32px
  title-lg:
    fontFamily: Inter
    fontSize: 20px
    fontWeight: '500'
    lineHeight: 28px
  body-lg:
    fontFamily: Inter
    fontSize: 18px
    fontWeight: '400'
    lineHeight: 28px
  body-md:
    fontFamily: Inter
    fontSize: 16px
    fontWeight: '400'
    lineHeight: 24px
  body-sm:
    fontFamily: Inter
    fontSize: 14px
    fontWeight: '400'
    lineHeight: 20px
  label-md:
    fontFamily: Inter
    fontSize: 12px
    fontWeight: '600'
    lineHeight: 16px
    letterSpacing: 0.05em
rounded:
  sm: 0.25rem
  DEFAULT: 0.5rem
  md: 0.75rem
  lg: 1rem
  xl: 1.5rem
  full: 9999px
spacing:
  unit: 4px
  gutter: 24px
  margin-mobile: 16px
  margin-desktop: 64px
  container-max: 1280px
---

## Brand & Style

This design system is built for a **Circular Material Exchange Platform**, where industrial efficiency meets ecological responsibility. The brand personality is **Professional, Sustainable, and Tech-Forward**. It bridges the gap between rugged industrial logistics and clean digital innovation, evoking a sense of trust, transparency, and environmental stewardship.

The design style is **Corporate / Modern** with a lean toward **Minimalism**. It utilizes a structured grid to manage complex data while employing subtle tactile cues—like soft shadows and intentional roundedness—to make the industrial process feel approachable and modern. High contrast and generous whitespace ensure that material specifications and exchange data are the primary focus, eliminating cognitive load for professional users.

## Colors

The palette is rooted in an "Eco-Industrial" aesthetic. 

- **Primary (Leaf Green):** Used for primary actions, success states, and elements representing sustainability and growth.
- **Secondary (Ocean Blue):** Reserved for institutional trust, navigation, and high-level categorization.
- **Exchange Orange:** Used as a functional accent to highlight transactional actions or movement of materials.
- **Neutrals:** A range of professional grays anchored by a deep "Earth Gray" for typography, ensuring WCAG AA compliance for contrast. 

The system defaults to a **light mode** to maintain a clean, "fresh" feeling, using pure white surfaces to symbolize the clarity and transparency of the circular economy.

## Typography

The design system utilizes **Inter** across all levels to ensure a systematic, utilitarian, and highly legible experience. 

- **Headlines:** Use tighter letter spacing and heavier weights to create a strong visual anchor for page sections.
- **Body Text:** Set at a comfortable 16px default with a 1.5x line height (24px) to ensure long-form technical specifications are readable.
- **Labels:** Uppercase styling is applied to `label-md` for metadata and technical specs, improving the "industrial" feel of the platform's data points.
- **Responsive Scaling:** On mobile devices, large display styles scale down to prevent excessive word-breaking and maintain visual harmony.

## Layout & Spacing

This design system uses a **Fixed Grid** model for desktop to ensure data density remains controlled and professional, while transitioning to a **Fluid Grid** for mobile devices.

- **Grid System:** A 12-column grid on desktop with 24px gutters. Elements should align to the grid to maintain an organized, "ledger-like" appearance suitable for material tracking.
- **Spacing Rhythm:** Based on a 4px baseline. All padding and margins should be increments of 4 (e.g., 8, 16, 24, 32, 48, 64).
- **Desktop:** 64px outer margins with a max-width container of 1280px.
- **Tablet:** 8-column grid with 32px margins.
- **Mobile:** 4-column grid with 16px margins. Content flows vertically; complex data tables should transition to card-based layouts.

## Elevation & Depth

To balance the industrial nature of the platform with modern tech-forwardness, the system uses **Tonal Layers** combined with **Ambient Shadows**.

- **Surface Levels:** The primary background is the lowest level. Content sits on "Surface" containers (White) which use a very subtle, diffused shadow (0px 2px 8px, 5% opacity black) to suggest a slight lift.
- **Interactive States:** Hovering over cards or buttons increases the shadow spread and decreases opacity slightly, providing a tactile "pull" effect.
- **Dividers:** Used sparingly. Instead of heavy lines, use 1px borders in a light neutral (`#E5E5E5`) or subtle shifts in background tone to define regions.

## Shapes

The shape language is defined as **Rounded**, striking a balance between "Industrial Square" and "Consumer Soft."

- **Standard Elements:** Buttons, input fields, and small cards use a 0.5rem (8px) radius.
- **Containers:** Large dashboard sections and main content areas use `rounded-lg` (16px) to soften the overall interface.
- **Data Tags:** Status chips and category labels use `rounded-xl` (24px) or full pill shapes to distinguish them from actionable buttons.

## Components

- **Buttons:** Primary buttons use a solid Leaf Green fill with white text. Secondary buttons use an Ocean Blue outline with a 1px stroke. All buttons have an 8px radius and a subtle lift on hover.
- **Input Fields:** Use a light gray background and a 1px border. On focus, the border transitions to Ocean Blue with a soft blue outer glow. Labels sit above the field in `label-md` bold.
- **Cards:** White background, 16px padding, and 16px corner radius. Used for material listings. They feature a 1px light stroke instead of a heavy shadow to maintain a clean, modern aesthetic.
- **Chips/Badges:** Used for material types (e.g., "Plastic," "Metal") and status (e.g., "Available," "Pending"). They use low-saturation versions of the brand colors with high-contrast text.
- **Lists:** High-density data lists use alternating row colors (White and a very faint Gray) to guide the eye without the need for heavy borders.
- **Checkboxes/Radios:** Use the Primary Green for the "selected" state, ensuring the platform's "sustainable" core is reinforced in every interaction.
- **Progress Indicators:** Use the "Sky Info" blue for active states in multi-step exchange workflows.