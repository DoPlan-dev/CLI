import { CSSProperties, HTMLAttributes, PropsWithChildren, ReactElement } from "./vendor/react.js";
import { PointerHitAreaMargins } from "./PanelResizeHandleRegistry.js";
export type PanelResizeHandleOnDragging = (isDragging: boolean) => void;
export type ResizeHandlerState = "drag" | "hover" | "inactive";
export type PanelResizeHandleProps = Omit<HTMLAttributes<keyof HTMLElementTagNameMap>, "id"> & PropsWithChildren<{
    className?: string;
    disabled?: boolean;
    hitAreaMargins?: PointerHitAreaMargins;
    id?: string | null;
    onDragging?: PanelResizeHandleOnDragging;
    style?: CSSProperties;
    tabIndex?: number;
    tagName?: keyof HTMLElementTagNameMap;
}>;
export declare function PanelResizeHandle({ children, className: classNameFromProps, disabled, hitAreaMargins, id: idFromProps, onDragging, style: styleFromProps, tabIndex, tagName: Type, ...rest }: PanelResizeHandleProps): ReactElement;
export declare namespace PanelResizeHandle {
    var displayName: string;
}
